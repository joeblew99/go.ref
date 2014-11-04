package namespace

import (
	"container/list"
	"io"
	"strings"

	"veyron.io/veyron/veyron/lib/glob"

	"veyron.io/veyron/veyron2/context"
	"veyron.io/veyron/veyron2/naming"
	"veyron.io/veyron/veyron2/services/mounttable/types"
	verror "veyron.io/veyron/veyron2/verror2"
	"veyron.io/veyron/veyron2/vlog"
)

const mountTableGlobReplyStreamLength = 100

type queuedEntry struct {
	me    *naming.MountEntry
	depth int // number of mount tables traversed recursively
}

// globAtServer performs a Glob at a single server and adds any results to the list.  Paramters are:
//   server    the server to perform the glob at.  This may include multiple names for different
//             instances of the same server.
//   pelems    the pattern to match relative to the mounted subtree.
//   l         the list to add results to.
//   recursive true to continue below the matched pattern
func (ns *namespace) globAtServer(ctx context.T, qe *queuedEntry, pattern *glob.Glob, l *list.List) error {
	server := qe.me
	pstr := pattern.String()
	vlog.VI(2).Infof("globAtServer(%v, %v)", *server, pstr)

	var lastErr error
	// Trying each instance till we get one that works.
	for _, s := range server.Servers {
		// If the pattern is finished (so we're only querying about the root on the
		// remote server) and the server is not another MT, then we needn't send the
		// query on since we know the server will not supply a new address for the
		// current name.
		if pattern.Finished() {
			_, n := naming.SplitAddressName(s.Server)
			if strings.HasPrefix(n, "//") {
				return nil
			}
		}

		// Don't further resolve s.Server.
		s := naming.MakeTerminal(s.Server)
		callCtx, _ := ctx.WithTimeout(callTimeout)
		client := ns.rt.Client()
		call, err := client.StartCall(callCtx, s, "Glob", []interface{}{pstr})
		if err != nil {
			lastErr = err
			continue // try another instance
		}

		// At this point we're commited to a server since it answered tha call.
		for {
			var e types.MountEntry
			err := call.Recv(&e)
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			// Prefix the results with the path of the mount point.
			e.Name = naming.Join(server.Name, e.Name)

			// Convert to the ever so slightly different name.MountTable version of a MountEntry
			// and add it to the list.
			x := &queuedEntry{
				me: &naming.MountEntry{
					Name:    e.Name,
					Servers: convertServers(e.Servers),
				},
				depth: qe.depth,
			}
			// x.depth is the number of severs we've walked through since we've gone
			// recursive (i.e. with pattern length of 0).
			if pattern.Len() == 0 {
				if x.depth++; x.depth > ns.maxRecursiveGlobDepth {
					continue
				}
			}
			l.PushBack(x)
		}

		var globerr error
		if err := call.Finish(&globerr); err != nil {
			return err
		}
		return globerr
	}

	return lastErr
}

// Glob implements naming.MountTable.Glob.
func (ns *namespace) Glob(ctx context.T, pattern string) (chan naming.MountEntry, error) {
	defer vlog.LogCall()()
	root, globPattern := naming.SplitAddressName(pattern)
	g, err := glob.Parse(globPattern)
	if err != nil {
		return nil, err
	}

	// Add constant components of pattern to the servers' addresses and
	// to the prefix.
	//var prefixElements []string
	//prefixElements, g = g.SplitFixedPrefix()
	//prefix := strings.Join(prefixElements, "/")
	prefix := ""
	if len(root) != 0 {
		prefix = naming.JoinAddressName(root, prefix)
	}

	// Start a thread to get the results and return the reply channel to the caller.
	servers := ns.rootName(prefix)
	if len(servers) == 0 {
		return nil, verror.Make(naming.ErrNoMountTable, ctx)
	}
	reply := make(chan naming.MountEntry, 100)
	go ns.globLoop(ctx, servers, prefix, g, reply)
	return reply, nil
}

// depth returns the directory depth of a given name.
func depth(name string) int {
	name = strings.Trim(name, "/")
	if name == "" {
		return 0
	}
	return strings.Count(name, "/") - strings.Count(name, "//") + 1
}

func (ns *namespace) globLoop(ctx context.T, servers []string, prefix string, pattern *glob.Glob, reply chan naming.MountEntry) {
	defer close(reply)

	// As we encounter new mount tables while traversing the Glob, we add them to the list 'l'.  The loop below
	// traverses this list removing a mount table each time and calling globAtServer to perform a glob at that
	// server.  globAtServer will send on 'reply' any terminal entries that match the glob and add any new mount
	// tables to be traversed to the list 'l'.
	l := list.New()
	l.PushBack(&queuedEntry{me: &naming.MountEntry{Name: "", Servers: convertStringsToServers(servers)}})
	atRoot := true

	// Perform a breadth first search of the name graph.
	for le := l.Front(); le != nil; le = l.Front() {
		l.Remove(le)
		e := le.Value.(*queuedEntry)

		// Get the pattern elements below the current path.
		suffix := pattern.Split(depth(e.me.Name))

		// Perform a glob at the server.
		err := ns.globAtServer(ctx, e, suffix, l)

		// We want to output this entry if:
		// 1. There was a real error, we return whatever name gave us the error.
		if err != nil && !notAnMT(err) {
			x := *e.me
			x.Name = naming.Join(prefix, x.Name)
			x.Error = err
			reply <- x
		}

		// 2. The current name fullfills the pattern and further servers did not respond
		//    with "".  That is, we want to prefer foo/ over foo.
		if suffix.Len() == 0 && !atRoot {
			x := *e.me
			x.Name = naming.Join(prefix, x.Name)
			reply <- x
		}
		atRoot = false
	}
}

var app = app || {};

(function() {
  'use strict';

  var Lists = new app.Collection('lists');
  var Todos = new app.Collection('todos');

  app.Lists = Lists;
  app.Todos = Todos;

  // Copied from meteor/todos/server/bootstrap.js.
  var data = [
    {name: 'Meteor Principles',
     contents: [
       ['Data on the Wire', 'Simplicity', 'Better UX', 'Fun'],
       ['One Language', 'Simplicity', 'Fun'],
       ['Database Everywhere', 'Simplicity'],
       ['Latency Compensation', 'Better UX'],
       ['Full Stack Reactivity', 'Better UX', 'Fun'],
       ['Embrace the Ecosystem', 'Fun'],
       ['Simplicity Equals Productivity', 'Simplicity', 'Fun']
     ]
    },
    {name: 'Languages',
     contents: [
       ['Lisp', 'GC'],
       ['C', 'Linked'],
       ['C++', 'Objects', 'Linked'],
       ['Python', 'GC', 'Objects'],
       ['Ruby', 'GC', 'Objects'],
       ['JavaScript', 'GC', 'Objects'],
       ['Scala', 'GC', 'Objects'],
       ['Erlang', 'GC'],
       ['6502 Assembly', 'Linked']
     ]
    },
    {name: 'Favorite Scientists',
     contents: [
       ['Ada Lovelace', 'Computer Science'],
       ['Grace Hopper', 'Computer Science'],
       ['Marie Curie', 'Physics', 'Chemistry'],
       ['Carl Friedrich Gauss', 'Math', 'Physics'],
       ['Nikola Tesla', 'Physics'],
       ['Claude Shannon', 'Math', 'Computer Science']
     ]
    }
  ];

  var timestamp = (new Date()).getTime();
  for (var i = 0; i < data.length; i++) {
    var listId = Lists.insert({name: data[i].name});
    for (var j = 0; j < data[i].contents.length; j++) {
      var info = data[i].contents[j];
      Todos.insert({listId: listId,
                    text: info[0],
                    done: false,
                    timestamp: timestamp,
                    tags: info.slice(1)});
      timestamp += 1;  // ensure unique timestamp
    }
  }
}());

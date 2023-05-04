let res = [
  db.users.drop(),
  db.apps.drop(),
  db.users.insertOne({
    email: "mabaums@gmail.com",
    name: "Mark Baumstein",
    isAdmin: true
  }),
  db.apps.insertOne({
    name: "app1",
    users: ["mabaums@gmail.com"],
    host: "app1",
    port: 22
  }),
  db.apps.insertOne({
    name: "app2",
    users: ["mabaums@gmail.com", "mbaumste@purdue.edu"],
    host: "app2",
    port: 22
  }),
]

printjson(res)

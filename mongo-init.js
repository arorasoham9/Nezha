let res = [
  db.users.drop(),
  db.apps.drop(),
  db.users.insertOne({email: "mabaums@gmail.com"}),
  db.users.insertOne({email: "mbaumste@purdue.edu"}),
  db.apps.insertOne({email: "mabaums@gmail.com", apps: ["app1", "app2"]}),
  db.apps.insertOne({email: "mbaumste@purdue.edu", apps: ["app2", "app3"]}),
]

printjson(res)

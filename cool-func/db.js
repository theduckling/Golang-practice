db.createUser({
  user: "maryam",
  pwd: "N1W8O+V93SPR7nWd6Ntjw",
  roles: [{ role: "readWrite", db: "train_statistics" }]
})


db.updateUser("maryam", {
  roles: [{ role: "read", db: "train_statistics" }]
})
db = db.getSiblingDB('coffeedb')

db.users.insert({
    name: "Fajrul Aulia",
    username: "fajrulaulia",
    email: "auliafajrul7@gmail.com",
    phonenumber: "082160100446",
    password: "$2y$12$tNKsfjSrVU/SHfdc/yNBP.YDogpFbCye9qpbESUPtYlQV1LgVCk9S",
    is_admin: true,
    is_verified: false,
    created_at: new Date(Date.now()),
    updated_at: new Date(Date.now())
});

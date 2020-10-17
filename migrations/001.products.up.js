db = db.getSiblingDB('coffeedb')

for (let i = 0; i <= 100; i++) {
    db.products.insert({
        name: "Kopi Bubuk Arabika " + i,
        description: "Bubuk kopi arabika yang terbuat dari kopi asli",
        sell_price: 40000,
        buy_price: 38000,
        created_at: new Date(Date.now()),
        updated_at: new Date(Date.now())
    });

    db.products.insert({
        name: "Kopi Bubuk Gayo " + i,
        description: "Bubuk kopi Gayo yang terbuat dari kopi asli",
        sell_price: 83000,
        buy_price: 80000,
        created_at: new Date(Date.now()),
        updated_at: new Date(Date.now())
    });

    db.products.insert({
        name: "Teh Pucuk Premium " + i,
        description: "Teh pucuk hijau asli dari Jawa Barat",
        sell_price: 45000,
        buy_price: 43000,
        created_at: new Date(Date.now()),
        updated_at: new Date(Date.now())
    });

    db.products.insert({
        name: "Gula Cap Payung " + i,
        description: "Gula asli dari tebu pilihan",
        sell_price: 12000,
        buy_price: 10000,
        created_at: new Date(Date.now()),
        updated_at: new Date(Date.now())
    });
}

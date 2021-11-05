# Task Introduction NoSQL and MongoDB 🍃

### Objektif
- Mampu untuk menggunakan Basic Operation MongoDB.
- Mampu menggunakan agregasi di Mongo DB

<br>

### Create
- [x] Insert 5 operators pada table operators.
- [x] Insert 3 product type.
- [x] Insert 2 product dengan product type id = 1, dan operators id = 3.
- [x] Insert 3 product dengan product type id = 2, dan operators id = 1.
- [x] Insert 3 product dengan product type id = 3, dan operators id = 4.
- [x] Insert product description pada setiap product.
- [x] Insert 3 payment methods.
- [x] Insert 5 user pada tabel user.
- [x] Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
- [x] Insert 3 product di masing-masing transaksi.

```
db.operators.insertMany([
  { "id": 1, "name": "telkomsel" },
  { "id": 2, "name": "xl" },
  { "id": 3, "name": "indosat" },
  { "id": 4, "name": "axis" },
  { "id": 5, "name": "tri" }
]);

db.productTypes.insertMany([
  { "id": 1, "name": "pulsa" },
  { "id": 2, "name": "token" },
  { "id": 3, "name": "internet" }
]);

db.products.insertMany([
  {
    "id": 1,
    "productTypeId": 1,
    "operatorId": 3,
    "name": "seru abiez"
  },
  {
    "id": 2,
    "productTypeId": 1,
    "operatorId": 3,
    "name": "paket sms hemat"
  },
])
db.products.insertMany([
  {
    "id": 3,
    "productTypeId": 2,
    "operatorId": 1,
    "name": "listrik rumah"
  },
  {
    "id": 4,
    "productTypeId": 2,
    "operatorId": 1,
    "name": "air PAM rumah"
  },
  {
    "id": 5,
    "productTypeId": 2,
    "operatorId": 1,
    "name": "pajak bangunan"
  }
])
db.products.insertMany([
  {
    "id": 6,
    "productTypeId": 3,
    "operatorId": 4,
    "name": "internet ngebut"
  },
  {
    "id": 7,
    "productTypeId": 3,
    "operatorId": 4,
    "name": "sosmed sikad"
  },
  {
    "id": 8,
    "productTypeId": 3,
    "operatorId": 4,
    "name": "stream malem"
  }
])

db.products.updateOne({ id: 1 }, { $set: { description: "paket buat nelpon irit bgt" } });
db.products.updateOne({ id: 2 }, { $set: { description: "sms seharian gratissssssss" } });
db.products.updateOne({ id: 3 }, { $set: { description: "buat bayar listrik rumah" } });
db.products.updateOne({ id: 4 }, { $set: { description: "bayar air ngga ribet" } });
db.products.updateOne({ id: 5 }, { $set: { description: "pajak dulu boss" } });
db.products.updateOne({ id: 6 }, { $set: { description: "internetan super ngebudzzzz" } });
db.products.updateOne({ id: 7 }, { $set: { description: "pansos di sosmed gass" } });
db.products.updateOne({ id: 8 }, { $set: { description: "stream pagi siang malam" } });

db.paymentMethods.insertMany([
  { "id": 1, "name": "debit" },
  { "id": 2, "name": "kredit" },
  { "id": 3, "name": "paylater" }
])

db.users.insertMany([
  {
    "id": 1,
    "name": "Hernowo Ari",
    "gender": "L",
    "dob": "2001-10-10"
  },
  {
    "id": 2,
    "name": "Erika Hana",
    "gender": "P",
    "dob": "2002-03-14"
  },
  {
    "id": 3,
    "name": "Hasna Nur",
    "gender": "P",
    "dob": "2001-08-29"
  },
  {
    "id": 4,
    "name": "Muhammad Raihan",
    "gender": "L",
    "dob": "1999-11-21"
  },
  {
    "id": 5,
    "name": "Hanif Edma",
    "gender": "L",
    "dob": "2002-02-12"
  }
])

db.transactions.insertMany([
  {
    "id": 1,
    "user_id": 1,
    "payment_method_id": 1,
    "total_qty": 3
  },
  {
    "id": 2,
    "user_id": 1,
    "payment_method_id": 1,
    "total_qty": 3
  },
  {
    "id": 3,
    "user_id": 1,
    "payment_method_id": 2,
    "total_qty": 3
  },
  {
    "id": 4,
    "user_id": 2,
    "payment_method_id": 3,
    "total_qty": 3
  },
  {
    "id": 5,
    "user_id": 2,
    "payment_method_id": 2,
    "total_qty": 3
  },
  {
    "id": 6,
    "user_id": 2,
    "payment_method_id": 2,
    "total_qty": 3
  },
  {
    "id": 7,
    "user_id": 3,
    "payment_method_id": 2,
    "total_qty": 3
  },
  {
    "id": 8,
    "user_id": 3,
    "payment_method_id": 2,
    "total_qty": 3
  }
])

db.transactions.updateOne({ id: 1 }, {
  $set: {
    products: [{
      "product_id": 1,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 3,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 6,
      "price": null,
      "qty": 1
    }]
  }
});
db.transactions.updateOne({ id: 2 }, {
  $set: {
    products: [{
      "product_id": 3,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 2,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 8,
      "price": null,
      "qty": 1
    }]
  }
});
db.transactions.updateOne({ id: 3 }, {
  $set: {
    products: [{
      "product_id": 4,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 7,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 2,
      "price": null,
      "qty": 1
    }]
  }
});
db.transactions.updateOne({ id: 4 }, {
  $set: {
    products: [{
      "product_id": 2,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 1,
      "price": null,
      "qty": 3
    },
    {
      "product_id": 5,
      "price": null,
      "qty": 1
    }]
  }
});
db.transactions.updateOne({ id: 5 }, {
  $set: {
    products: [{
      "product_id": 5,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 7,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 8,
      "price": null,
      "qty": 1
    }]
  }
});
db.transactions.updateOne({ id: 6 }, {
  $set: {
    products: [{
      "product_id": 1,
      "price": null,
      "qty": 3
    },
    {
      "product_id": 4,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 7,
      "price": null,
      "qty": 1
    }]
  }
});
db.transactions.updateOne({ id: 7 }, {
  $set: {
    products: [{
      "product_id": 6,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 2,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 8,
      "price": null,
      "qty": 1
    }]
  }
});
db.transactions.updateOne({ id: 8 }, {
  $set: {
    products: [{
      "product_id": 7,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 5,
      "price": null,
      "qty": 1
    },
    {
      "product_id": 3,
      "price": null,
      "qty": 1
    }]
  }
});
```

<br>

### Read
- [x] Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
- [x] Tampilkan product dengan id = 3.
- [x] Hitung jumlah user / pelanggan dengan status gender Perempuan.
- [x] Tampilkan data pelanggan dengan urutan sesuai nama abjad.
- [x] Tampilkan 5 data product.

```
db.users.find({ gender: 'L' })
db.products.find({ id: 3 })
db.users.find({ gender: 'P' }).count()
db.users.find().sort({ name: 1 })
db.products.find().limit(5)
```

<br>

### Update
- [x] Ubah data product id 1 dengan nama ‘product dummy’.
- [x] Ubah qty = 3 pada transaction detail dengan product id 1.

```
db.products.updateOne({ id: 1 }, { $set: { name: "product dummy" } })

// https://stackoverflow.com/questions/11372065/mongodb-how-do-i-update-a-single-subelement-in-an-array-referenced-by-the-inde
db.transactions.updateMany({ 'products.product_id': 1 }, { $set: { 'products.$.qty': 3 } })
```

### Delete
- [x] Delete data pada tabel product dengan id 1.
- [x] Delete pada pada tabel product dengan product type id 1.

```
db.products.deleteOne({ id: 1 })
db.products.deleteMany({ productTypeId: 1 })
```
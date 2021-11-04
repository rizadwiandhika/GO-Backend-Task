/* 
Insert:
  Insert 5 operators pada table operators.
  Insert 3 product type.
  Insert 2 product dengan product type id = 1, dan operators id = 3.
  Insert 3 product dengan product type id = 2, dan operators id = 1.
  Insert 3 product dengan product type id = 3, dan operators id = 4.
  Insert product description pada setiap product.
  Insert 3 payment methods.
  Insert 5 user pada tabel user.
  Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)
  Insert 3 product di masing-masing transaksi.
*/
INSERT INTO operators (name) VALUES ("telkomsel"), ("xl"), ("indosat"), ("axis"), ("tri");
INSERT INTO product_types (name) VALUES ("pulsa"), ("token"), ("internet");
INSERT INTO products (product_type_id, operator_id, name) VALUES (1, 3, "paket nelpon hemat"), (1, 3, "paket sms hemat");
INSERT INTO products (product_type_id, operator_id, name) VALUES 
	(2, 1, "listrik rumah"), 
  (2, 1, "air PAM rumah"),
  (2, 1, "pajak bangunan");
INSERT INTO products (product_type_id, operator_id, name) VALUES 
	(3, 4, "internet ngebut"), 
  (3, 4, "sosmed sikad"),
  (3, 4, "stream malem");
INSERT INTO product_descriptions (product_id, description) VALUES 
	(1, "paket buat nelpon irit bgt"),
  (2, "sms seharian gratissssssss"),
  (3, "buat bayar listrik rumah"),
  (4, "bayar air ngga ribet"),
  (5, "pajak dulu boss"),
  (6, "internetan super ngebudzzzz"),
  (7, "pansos di sosmed gass"),
  (8, "stream pagi siang malam");
INSERT INTO payment_methods (name) VALUES ("debit"), ("kredit"), ("paylater");
INSERT INTO users (name, gender, dob) VALUES 
	("Hernowo Ari", 'L', "2001-10-10"), 
  ("Erika Hana", 'P', "2002-03-14"),
  ("Hasna Nur", 'P', "2001-08-29"),
  ("Muhammad Raihan", 'L', "1999-11-21"),
  ("Hanif Edma", 'L', "2002-02-12");
INSERT INTO transactions (user_id, payment_method_id, total_qty) VALUES 
  (1, 1, 3),
  (1, 1, 3),
  (1, 2, 3),
  (2, 3, 3),
  (2, 2, 3),
  (2, 2, 3),
  (3, 2, 3),
  (3, 2, 3),
  (3, 1, 3),
  (4, 2, 3),
  (4, 1, 3),
  (4, 3, 3),
  (5, 3, 3);
  (5, 1, 3);
  (5, 3, 3);
INSERT INTO detail_transactions (transaction_id, product_id, qty) VALUES 
  (1, 1, 1),
  (1, 3, 1),
  (1, 6, 1),
  (2, 3, 1),
  (2, 2, 1),
  (2, 8, 1),
  (3, 4, 1),
  (3, 7, 1),
  (3, 2, 1),
  (4, 2, 1),
  (4, 1, 1),
  (4, 5, 1),
  (5, 5, 1),
  (5, 7, 1),
  (5, 8, 1),
  (6, 1, 1),
  (6, 4, 1),
  (6, 7, 1),
  (7, 6, 1),
  (7, 2, 1),
  (7, 8, 1),
  (8, 7, 1),
  (8, 5, 1),
  (8, 3, 1);


/*
Select:
  Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
  Tampilkan product dengan id = 3.
  Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
  Hitung jumlah user / pelanggan dengan status gender Perempuan.
  Tampilkan data pelanggan dengan urutan sesuai nama abjad
  Tampilkan 5 data pada data product
*/
SELECT name from users where gender = 'L';
SELECT * FROM products WHERE id = 3;
SELECT * FROM users WHERE created_at - CURRENT;
SELECT * FROM users 
	WHERE name LIKE '%a%' AND TIMESTAMPDIFF(DAY, created_at, CURRENT_TIMESTAMP) < 7;
SELECT COUNT(*) AS "Total Perempuan" FROM users WHERE gender = 'P';
SELECT name from users ORDER BY name;
SELECT * FROM products LIMIT 5;


/*
Update:
    Ubah data product id 1 dengan nama ‘product dummy’.
    Update qty = 3 pada transaction detail dengan product id 1.
*/
UPDATE products SET name = 'product_dummy' WHERE id = 1;
UPDATE detail_transactions SET qty = 3 WHERE product_id = 1;


/* 
Join, Union, Sub-query, Function: 
    Gabungkan data transaksi dari user id 1 dan user id 2.
    Tampilkan jumlah harga transaksi user id 1.
    Tampilkan total transaksi dengan product type 2.
    Tampilkan semua field table product dan field name table product type yang saling berhubungan.
    Tampilkan semua field table transaction, field name table product dan field name table user.
    Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
    Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
    Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
*/

SELECT * FROM transactions WHERE user_id = 1
UNION
SELECT * FROM transactions WHERE user_id = 2;

SELECT SUM(total_price) AS "user 1 total price" FROM transactions WHERE user_id = 1;

SELECT SUM(dt.price) AS "product 2 total transactions" FROM detail_transactions dt 
JOIN products p ON (dt.product_id = p.id AND p.product_type_id = 2);

SELECT p.*, pt.* FROM products p 
JOIN product_types pt ON (p.product_type_id = pt.id);

SELECT t.*, p.name, u.name FROM transactions t 
JOIN users u ON (t.user_id = u.id)
JOIN detail_transactions dt ON (t.id = dt.transaction_id)
JOIN products p ON (dt.product_id = p.id)  
ORDER BY t.id;

DELIMITER $$
CREATE TRIGGER before_delete_trasaction BEFORE DELETE
ON transactions FOR EACH ROW
BEGIN
	DELETE FROM detail_transactions WHERE transaction_id = OLD.id;
END
$$
DELIMITER ;

DELIMITER $$
CREATE TRIGGER after_delete_detail_trasaction AFTER DELETE
ON detail_transactions FOR EACH ROW
BEGIN
	UPDATE transactions SET total_qty = total_qty + OLD.qty WHERE id = OLD.transaction_id;
END
$$
DELIMITER ;

SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM detail_transactions);
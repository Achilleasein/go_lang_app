Use wallet_db;

Create table `wallets` (id int, wallet varchar(20), money int, username varchar(20), password varchar(20));

Insert into wallets (id, wallet, money) values (1, 'wallet-1', 0);
Insert into wallets (id, wallet, money) values (2, 'wallet-2', 0);
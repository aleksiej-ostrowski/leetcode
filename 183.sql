Create table If Not Exists Customers (id int, name varchar(255))
Create table If Not Exists Orders (id int, customerId int)
Truncate table Customers
insert into Customers (id, name) values ('1', 'Joe')
insert into Customers (id, name) values ('2', 'Henry')
insert into Customers (id, name) values ('3', 'Sam')
insert into Customers (id, name) values ('4', 'Max')
Truncate table Orders
insert into Orders (id, customerId) values ('1', '3')
insert into Orders (id, customerId) values ('2', '1')


SELECT `t1`.`name` as 'Customers'
FROM `Customers` AS t1
LEFT OUTER JOIN `Orders` as t2 
ON `t1`.`id` = `t2`.`customerId`
WHERE `t2`.`customerId` is NULL

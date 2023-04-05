Create table If Not Exists Employee (id int, name varchar(255), salary int, departmentId int)
Create table If Not Exists Department (id int, name varchar(255))
Truncate table Employee
insert into Employee (id, name, salary, departmentId) values ('1', 'Joe', '70000', '1')
insert into Employee (id, name, salary, departmentId) values ('2', 'Jim', '90000', '1')
insert into Employee (id, name, salary, departmentId) values ('3', 'Henry', '80000', '2')
insert into Employee (id, name, salary, departmentId) values ('4', 'Sam', '60000', '2')
insert into Employee (id, name, salary, departmentId) values ('5', 'Max', '90000', '1')
Truncate table Department
insert into Department (id, name) values ('1', 'IT')
insert into Department (id, name) values ('2', 'Sales')


SELECT 
t2.`name` AS 'Department', 
t1.`name` AS 'Employee', 
t1.`salary` AS 'Salary' 
FROM `Employee` AS t1  
LEFT JOIN `Department` AS t2 
ON t1.`departmentId` = t2.`id`
INNER JOIN (
SELECT 
t4.`name` AS 'Department', 
MAX(t3.`salary`) AS 'Salary' 
FROM `Employee` AS t3  
LEFT JOIN `Department` AS t4 
ON t3.`departmentId` = t4.`id` 
GROUP BY t3.`departmentId`
) AS t5 
ON
t2.`name` = t5.`Department` AND
t1.`salary` = t5.`Salary`

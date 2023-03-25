Create table If Not Exists Employee (id int, salary int)
Truncate table Employee
insert into Employee (id, salary) values ('1', '100')
insert into Employee (id, salary) values ('2', '200')
insert into Employee (id, salary) values ('3', '300')


SELECT max(salary) SecondHighestSalary 
FROM Employee 
WHERE salary < (SELECT max(salary) FROM Employee)

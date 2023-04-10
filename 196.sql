Create table If Not Exists Person (Id int, Email varchar(255))
Truncate table Person
insert into Person (id, email) values ('1', 'john@example.com')
insert into Person (id, email) values ('2', 'bob@example.com')
insert into Person (id, email) values ('3', 'john@example.com')

/* https://leetcode.com/problems/delete-duplicate-emails/solutions/2627589/my-sql-solution */

DELETE p1 FROM Person AS p1, Person AS p2 
WHERE p1.`email`=p2.`email` AND p1.`id`>p2.`id`


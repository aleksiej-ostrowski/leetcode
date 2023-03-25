Create table If Not Exists Employee (Id int, Salary int)
Truncate table Employee
insert into Employee (id, salary) values ('1', '100')
insert into Employee (id, salary) values ('2', '200')
insert into Employee (id, salary) values ('3', '300')


CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  DECLARE OFS INT;
  SELECT N-1 INTO OFS;
  RETURN (
      SELECT DISTINCT salary FROM Employee 
      ORDER BY salary DESC LIMIT OFS, 1    
  );
END

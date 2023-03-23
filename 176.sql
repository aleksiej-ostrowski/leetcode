SELECT max(salary) SecondHighestSalary 
FROM Employee 
WHERE salary < (SELECT max(salary) FROM Employee)

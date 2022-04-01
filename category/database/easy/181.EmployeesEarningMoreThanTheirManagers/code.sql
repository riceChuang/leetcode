

SELECT e.name as Employee
FROM Employee as e Left Join Employee as en
on e.managerId = en.id
WHERE e.salary > en.salary



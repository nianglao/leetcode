CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
      select max(E1.Salary)
      from Employee as E1
      where (N)=(select count(distinct(E2.Salary))
      from Employee as E2
      where E2.Salary >= E1.Salary)
  );
END

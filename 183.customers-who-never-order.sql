# Write your MySQL query statement below

select distinct Customers.Name as `Customers`
from Customers
left join Orders on Customers.Id = Orders.CustomerId
where Orders.CustomerId is null
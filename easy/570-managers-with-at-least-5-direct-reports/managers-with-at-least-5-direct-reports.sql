select Employee.name 
from Employee 
join (
select managerId,name,count(id)  cont
from Employee
where managerId is not null 
group by managerId

) as aux
on aux.managerId = id
where aux.cont >= 5 
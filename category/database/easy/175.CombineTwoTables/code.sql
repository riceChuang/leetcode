


SELECT FirstName, LastName, City, State
FROM Person Left Join Address
on Person.personID = Address.personID;
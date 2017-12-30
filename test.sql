create database GOTEST;
use GOTEST

create table OperationsAllowed (id varchar(3), error int, host int, op varchar(12));
insert into OperationsAllowed (id, error, host, op) values ('1', 0, 1, 'purchase');
insert into OperationsAllowed (id, error, host, op) values ('1', 0, 1, 'sale');
insert into OperationsAllowed (id, error, host, op) values ('1', 0, 1, 'generic');

/*only if you want to check what you have created*/
select * from OperationsAllowed;

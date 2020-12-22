# mysql-benchmarker
a benchmarker that benchmark how many connections can mySQL accept before it crashes

This application benchmarks the performance of a database in terms of how many concurrent connections it can receive

This application requires a MySQL or MySQL binary compatible database that have a database that have a table called test. you need to provide the application with a golang standard MySQL connection string to be able to connect to it

This application listens on port 6443 for a POST request on / that asks for a form value of *sqlconnectionstring*
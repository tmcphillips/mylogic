## 2019-06-10 Think about goals for the MyLogic project

### Some observations
The following are observations I have made while using logic programming in the course of research:

- It has been extremely useful to **model workflow execution traces and provenance** as Datalog facts and to **query them** using Datalog (or tabled Prolog using XSB).

- Datalog itself is not a complete programming language, one reason that using XSB is so helpful.  

- For example, an XSB program can perform Datalog-style queries and then employ Prolog features to output the results either as formatted reports or as Graphviz visualizations.

- However, XSB itself is has limitations as an application programming environment, and as an environment for building code libraries for use in applications.

- In some ways I can be **much more productive**--and agile--developing data models and implementing queries **using logic programming** than I can using database management systems and SQL.

- There are many variations on logic programming language capabilities in different implementations, each with their own advantages and disadvantages.

- Understanding how all the various logic programming paradigms and features actually work, and how best to take advantage of each is difficult.

- It is easiest to understand a framework if you are its developer, and to understand a paradigm fully if you have authored an implementation of that paradigm.

- It is easiest to write code for and debug code if you understand how the programming language and its runtime work.

### Vision

Here are some desires, hopes and long-term goals for the MyLogic project:
- Make the **logic programming paradigm a practical alternative** to full-blown relational databases for application development.  Instead of embedding SQLite in an application, I'd like to **define an extensional database in plain text files and work with that data using Datalog-style queries**.

- Make it easier for everyone to take advantage of logic programming and the capabilities it brings over, say, SQL or SPARQL, and to apply these approaches to **making scientific workflows and data provenance more transparent**, and **research results more reproducible**.

-  Develop an **approach for enabling *research transparency* and *data transparency* that is itself *transparent***.  Why trust obscure logic programs and tools that someone claims make research easier to understand and the provenance of data easier to visualize and query, when the very tools and approaches for achieving that transparency are themselves utterly opaque?  






 

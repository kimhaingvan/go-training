# Code challenge  
  
Using the provided data (tickets.json and users.json and organization.json)
Write a simple command live application to search the data and return the results in a human readable format.

- Feel free to use libraries or roll your own code as you see fit.
- Where the data exists, value from any related entities MUST be included in the results, i.e.
  - Searching **organization MUST** returnn its **ticket subject** and **users name**.
  - Searching **users MUST** return his/ber **assignee ticket subject** and **submitted ticket subject** annd his/her **organization name**.
  - Searching **tickets MUST** return its **assignee nname, submitter name**, and **organizationn name**.
- The user should be able to search on any field, full value matching is fine (e.g. "mar" won't return "mary").
- The user should also be able to search for empty values, e.g. where description is empty.
- Search can get pretty complicated pretty easily, we jsut want to see that you can code a basic application.

## SAMPLE

### Search

![alt text](https://github.com/builamquangngoc91/go-training/blob/main/assets/1.png)

![alt text](https://github.com/builamquangngoc91/go-training/blob/main/assets/2.png)

### No result

![alt text](https://github.com/builamquangngoc91/go-training/blob/main/assets/3.png)

### List of searchable fields

![alt text](https://github.com/builamquangngoc91/go-training/blob/main/assets/4.png)


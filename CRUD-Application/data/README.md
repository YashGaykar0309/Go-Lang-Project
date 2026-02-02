# Adding data to database

There are several "easier" ways to do this, but this is the most repeatable for different levels of experience as well as different operating systems.

1. Start database - ./postgres_start.sh
   - note: you can make this persistent by specifying volumes in the script such as adding:
   ```
   -e PGDATA=/var/lib/postgresql/data/pgdata \
   -v /Users/frankmoley/.local/docker/data:/var/lib/postgresql/data \
   ```
   
2. Exec and launch psql into docker container
   ```
   docker exec -it postgres-db psql -U admin -d mydb
   ```

3. Copy/paste schema file and then data file from this directory into psql



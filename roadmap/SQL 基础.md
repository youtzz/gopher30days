## SQL 基础

### SQL Update

```sql
UPDATE table_name
SET column1 = value1, column2 = value2, ...
WHERE confiion;
```

### SQL Delete

```sql
DELETE FROM table_name WHERE confition;
```

### SQL Min and Max

The `MIN()` function returns the smallest value of the selected column.

```sql
SELECT MIN(column_name)
FROM table_name
WHERE condition;
```

The `MAX()` function returns the largest value of the selected column.

```sq
SELECT MAX(column_name)
FROM table_name
WHERE condition;
```

### SQL Count, Avg, Sum

The `COUNT()` function returns the number of rows that matches a specified criterion.

```sql
SELECT COUNT(column_name)
FROM table_name
WHERE condition;
```

The `AVG()` function returns the average value of a numeric column. 

```sql
SELECT AVG(column_name)
FROM table_name
WHERE condition;
```

The `SUM()` function returns the total sum of a numeric column. 

```sql
SELECT SUM(column_name)
FROM table_name
WHERE condition;
```


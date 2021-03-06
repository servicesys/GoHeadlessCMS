-- CREATE SCHEMA headless_cms;

SET search_path TO headless_cms;

-- 
-- This script is generated automatically and should not be modified 
-- Changes should only be made through DbDesigner 
-- 

DROP TABLE CONTENT_CATEGORY CASCADE; 


DROP TABLE CONTENT_VALUE CASCADE; 


DROP TABLE CONTENT_TYPE CASCADE; 



 
CREATE TABLE CONTENT_CATEGORY ( 
  COD CHAR(40) NOT NULL, 
  DESCRIPTION VARCHAR(120) NOT NULL 
); 
ALTER TABLE CONTENT_CATEGORY ADD CONSTRAINT CONTENT_CATEGORY_PK PRIMARY KEY (COD); 
 


 
CREATE TABLE CONTENT_VALUE ( 
  UUID UUID NOT NULL, 
  CONTENT_CATEGORY_COD CHAR(40) NOT NULL, 
  CONTENT_TYPE_COD CHAR(40) NOT NULL, 
  VALUE JSONB NOT NULL, 
  CREATED_ON TIMESTAMP NOT NULL, 
  MODIFIED_ON TIMESTAMP NOT NULL, 
  CONTENT_STATUS CHAR(10) NOT NULL, 
  NAME VARCHAR(120) NULL 
); 
ALTER TABLE CONTENT_VALUE ADD CONSTRAINT CONTENT_VALUE_PK PRIMARY KEY (UUID); 
 
CREATE INDEX CONTENT_VALUE_TYPE_COD_FK ON CONTENT_VALUE (CONTENT_TYPE_COD); 
 
CREATE INDEX CONTENT_VALUE_CATEGORY_COD_FK ON CONTENT_VALUE (CONTENT_CATEGORY_COD); 
 
CREATE UNIQUE INDEX CONTENT_VALUE_NAME_IDX ON CONTENT_VALUE (NAME); 


 
CREATE TABLE CONTENT_TYPE ( 
  COD CHAR(40) NOT NULL, 
  METADATA JSONB NOT NULL, 
  DESCRIPTION VARCHAR(120) NOT NULL 
); 
ALTER TABLE CONTENT_TYPE ADD CONSTRAINT CONTENT_TYPE_PK PRIMARY KEY (COD); 
 







 
ALTER TABLE CONTENT_VALUE ADD CONSTRAINT TYPE_VALUE_REL FOREIGN KEY (CONTENT_TYPE_COD) REFERENCES CONTENT_TYPE(COD);
 
ALTER TABLE CONTENT_VALUE ADD CONSTRAINT CATEGORY_VALUE_REL FOREIGN KEY (CONTENT_CATEGORY_COD) REFERENCES CONTENT_CATEGORY(COD);
 





























 
comment on column CONTENT_VALUE.CONTENT_STATUS is 'REM - REMOVIDO - PUB - PUBLICADO APR - APROVACAO';














































 

-- End of generated script 












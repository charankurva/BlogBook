ALTER TABLE blog ADD CONSTRAINT chk_author CHECK (author>=0);
ALTER TABLE blog ADD CONSTRAINT chk_category CHECK (categoryID>=0);

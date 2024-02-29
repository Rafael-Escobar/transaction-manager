BEGIN;

INSERT INTO operation_types (id,is_debit, description) VALUES 
    (1,true, 'COMPRA A VISTA'),
    (2,true, 'COMPRA PARCELADA'),
    (3,true, 'SAQUE'),
    (4,false, 'PAGAMENTO');

COMMIT;
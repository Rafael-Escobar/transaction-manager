BEGIN;

INSERT INTO operation_types (id, description) VALUES 
    (1, 'COMPRA A VISTA'),
    (2, 'COMPRA PARCELADA'),
    (3, 'SAQUE'),
    (4, 'PAGAMENTO');

COMMIT;
BEGIN;

DELETE FROM operation_types WHERE description IN (
    'COMPRA A VISTA',
    'COMPRA PARCELADA',
    'SAQUE',
    'PAGAMENTO');

COMMIT;
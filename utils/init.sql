
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    document_type VARCHAR(10) CHECK (document_type IN ('CPF', 'CNPJ')),
    document_number VARCHAR(20),
    address VARCHAR(20),
    state VARCHAR(2),
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,    
    account_type VARCHAR(20) NOT NULL CHECK (account_type IN ('CHECKING', 'SAVINGS')),    
    balance BIGINT NOT NULL DEFAULT 0,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    
    CONSTRAINT fk_account_user FOREIGN KEY (user_id) REFERENCES users(id)
);


CREATE TABLE transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    account_id UUID NOT NULL,
    type VARCHAR(20) NOT NULL CHECK (type IN ('DEPOSIT', 'WITHDRAWAL', 'TRANSFER', 'PAYMENT')),
    method VARCHAR(30) NOT NULL CHECK (method IN ('PIX', 'TED', 'CASH_DEPOSIT', 'CASH_WITHDRAWAL')),
    destination_value TEXT,
    status VARCHAR(20) NOT NULL CHECK (status IN ('PENDING', 'COMPLETED', 'FAILED')),
    amount BIGINT NOT NULL CHECK (amount > 0),
    currency CHAR(3) NOT NULL DEFAULT 'BRL',
    description TEXT,
    related_transaction_id UUID, 
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_account FOREIGN KEY (account_id) REFERENCES accounts(id),
    CONSTRAINT fk_related_transaction FOREIGN KEY (related_transaction_id) REFERENCES transactions(id)
);


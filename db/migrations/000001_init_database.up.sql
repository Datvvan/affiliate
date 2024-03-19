CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID DEFAULT Uuid_generate_v4() PRIMARY KEY,
    email varchar UNIQUE NOT NULL,
    type varchar(50),
    ref_code varchar,
    intermediary UUID,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE user_transaction (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    type VARCHAR(50),
    status VARCHAR(50),
    update_at TIMESTAMP DEFAULT NOW(),
    create_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE affiliate_referrals (
    id SERIAL PRIMARY KEY,
    affiliate UUID,
    referral UUID,
    is_conversion BOOLEAN DEFAULT '0' NOT NULL,
    is_canceled_sub BOOLEAN DEFAULT '0' NOT NULL,
    commission_status VARCHAR(50),
    transaction_id INT,
    batch_id VARCHAR,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (affiliate) REFERENCES users(id),
    FOREIGN KEY (referral) REFERENCES users(id),
    FOREIGN KEY (transaction_id) REFERENCES user_transaction(id)

);

CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    member_type VARCHAR(50),
    end_of_trial_time TIMESTAMP,
    transaction_id int,
    expired_time TIMESTAMP,
    update_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (transaction_id) REFERENCES user_transaction(id)
);
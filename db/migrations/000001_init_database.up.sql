CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID DEFAULT Uuid_generate_v4() PRIMARY KEY,
    email varchar UNIQUE NOT NULL,
    type varchar(50),
    ref_code varchar,
    commission_amount integer,
    intermediary UUID,
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE affiliate_referrals (
    id SERIAL PRIMARY KEY,
    affiliate UUID,
    referral UUID,
    is_conversion BOOLEAN DEFAULT '0',
    is_canceled_sub BOOLEAN DEFAULT '0',
    commission_status VARCHAR(50),
    create_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW(),

    FOREIGN KEY (affiliate) REFERENCES users(id),
    FOREIGN KEY (referral) REFERENCES users(id)
);

CREATE TABLE subscriptions (
    id SERIAL PRIMARY KEY,
    user_id UUID,
    member_type VARCHAR(50),
    last_paid_date TIMESTAMP,
    end_of_trial_time TIMESTAMP,
    expired_time TIMESTAMP,
    update_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
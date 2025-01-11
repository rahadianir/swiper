CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "username" VARCHAR NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL,
    "age" INTEGER NOT NULL,
    "gender" VARCHAR NOT NULL,
    "location" VARCHAR NOT NULL,
    "is_premium" BOOLEAN NOT NULL,
    "is_verified" BOOLEAN NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ
);

-- data seed
INSERT INTO users (name, username, password, age, gender, location, is_premium, is_verified, created_at) VALUES
('Rachel', 'Rachel_1', 'fd5587950e947330', 42, 'male', 'Los Angeles', false, false, now()),
('Quentin', 'Quentin_2', '820d12e1f7b983ca', 47, 'male', 'Denver', false, false, now()),
('Kevin', 'Kevin_3', 'd18bfefedadd8db2', 18, 'female', 'Los Angeles', true, true, now()),
('George', 'George_4', 'dd933e5c781f0cc6', 24, 'female', 'Houston', false, false, now()),
('Kevin', 'Kevin_5', '941ebc604f8829fd', 39, 'male', 'Houston', true, true, now()),
('Luna', 'Luna_6', '4bb34f653284ce53', 39, 'female', 'Denver', false, false, now()),
('David', 'David_7', '524aee78d514a7e8', 19, 'female', 'Chicago', false, false, now()),
('Charlie', 'Charlie_8', 'd6a0c0f4affa538c', 25, 'female', 'Denver', false, false, now()),
('Julia', 'Julia_9', '98af171a3fd90c14', 39, 'male', 'Chicago', false, false, now()),
('Rachel', 'Rachel_10', '9e7e09f68ed02012', 22, 'male', 'Seattle', false, false, now()),
('Oscar', 'Oscar_11', 'ba8d3dc3c11965c7', 44, 'female', 'Chicago', true, true, now()),
('Julia', 'Julia_12', '18750d5c07dc58c6', 47, 'female', 'Boston', true, true, now()),
('Luna', 'Luna_13', 'b11c88581163f161', 20, 'male', 'Chicago', false, false, now()),
('Nina', 'Nina_14', '045cd7dbc3c9b5ab', 20, 'male', 'Houston', false, false, now()),
('Nina', 'Nina_15', 'ae548249c881fad9', 32, 'female', 'Miami', false, false, now()),
('Bob', 'Bob_16', 'bc7d104a97230ed3', 20, 'male', 'Houston', true, true, now()),
('Isaac', 'Isaac_17', '481092ab86584541', 27, 'female', 'Denver', false, false, now()),
('Steve', 'Steve_18', '29846861753909ff', 35, 'male', 'Dallas', true, true, now()),
('Hannah', 'Hannah_19', '7403f23803e0554e', 28, 'female', 'Seattle', false, false, now()),
('Rachel', 'Rachel_20', '6086fead082b2a6f', 46, 'female', 'Seattle', true, true, now()),
('George', 'George_21', '65340f22ed582bbc', 43, 'female', 'Chicago', false, false, now()),
('Steve', 'Steve_22', 'c080dd8f0ce7538b', 34, 'male', 'Chicago', true, true, now());

INSERT INTO groups (name, isactive)
VALUES ('group1', true),
       ('group2', true),
       ('group3', true);

INSERT INTO users (name, password, email, isActive, group_id)
VALUES ('user1', '$2a$10$XQADx0xZrXlGC/StdpJivOrQhSspOR32dRaowygMmPqYyiBkgj/qO', 'user1@user.com', true, 1),
       ('user2', '$2a$10$XQADx0xZrXlGC/StdpJivOrQhSspOR32dRaowygMmPqYyiBkgj/qO', 'user2@user.com', true, 2),
       ('user3', '$2a$10$XQADx0xZrXlGC/StdpJivOrQhSspOR32dRaowygMmPqYyiBkgj/qO', 'user3@user.com', true, 3),
       ('user4', '$2a$10$XQADx0xZrXlGC/StdpJivOrQhSspOR32dRaowygMmPqYyiBkgj/qO', 'user4@user.com', true, 1);
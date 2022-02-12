create database aamauDB;

create table Ingredient (
    ingredientId INT(4) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    ingredientName VARCHAR(50) NOT NULL,
    ingredientWeight FLOAT(24) NOT NULL,
    ingredientAmount  int UNSIGNED NOT NULL
);

create table Cake (
    cakeId INT(4) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cakeName VARCHAR(50) NOT NULL,
    dayNeeded INT UNSIGNED NOT NULL,
    price FLOAT(24) NOT NULL
);

create table Recipe (
    recipeId INT(4) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cakeId INT(4) UNSIGNED NOT NULL,
    ingredientId INT(4) UNSIGNED NOT NULL,
    ingredientAmountRequired INT UNSIGNED NOT NULL 
);

create table Orders (
    orderId INT(4) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    orderDate DATE NOT NULL,
    deliveryDate DATE NOT NULL,
    userId INT(4) UNSIGNED NOT NULL, 
    cakeId INT(4) UNSIGNED NOT NULL, 
    amount INT UNSIGNED NOT NULL, 
    totalPrice FLOAT(24) NOT NULL
);

create table User (
    userId INT(4) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    userName VARCHAR(50) NOT NULL,
    contactNo VARCHAR(10) NOT NULL,
    email VARCHAR(100),
    deliveryAddress VARCHAR(100) NOT NULL
);
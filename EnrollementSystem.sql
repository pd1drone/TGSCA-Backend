CREATE DATABASE TGSCA;

USE TGSCA;

CREATE TABLE `Users` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `Username` varchar(255),
  `Password` varchar(255),
  `IsAdmin` bool,
  `PlainPassword` varchar(255)
);

CREATE TABLE `Students` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `StudentNumber` integer,
  `UserID` integer,
  `FirstName` varchar(255),
  `LastName` varchar(255),
  `MiddleName` varchar(255),
  `Email` varchar(255),
  `DateOfBirth` varchar(255),
  `GradeLevel` varchar(255),
  `ContactNumber` varchar(255),
  `Address` varchar(255)
);

CREATE TABLE `Requirements` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `StudentNumber` integer,
  `UploadedFile` longtext,
  `RequirementType` varchar(255)
);

CREATE TABLE `Subjects` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `Subject` varchar(255),
  `GradeLevel` varchar(255),
  `Schedule` varchar(255),
  `TeachersID` integer
);

CREATE TABLE `Enrolled` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `StudentNumber` integer,
  `SubjectID` integer
);

CREATE TABLE `EnrolledPending`(
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `StudentNumber` integer,
  `SubjectID` integer
)

CREATE TABLE `Enrollment` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `StudentNumber` integer,
  `ProgressCard` longtext,
  `ProgressCardStatus` varchar(255),
  `Form137` longtext,
  `Form137Status` varchar(255),
  `GoodMoral` longtext,
  `GoodMoralStatus` varchar(255),
  `RegistrationFee` longtext,
  `RegistrationFeeStatus` varchar(255),
  `EnrollmentStatus` varchar(255)
);

CREATE TABLE `Appointments` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `Name` varchar(255),
  `Email` varchar(255),
  `ContactNumber` varchar(255),
  `StudentNumber` integer,
  `AppointmentType` varchar(255),
  `AppointmentDescription` longtext,
  `AppointmentDate` varchar(255)
);

CREATE TABLE `Teachers` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `TeacherName` varchar(255)
);


INSERT INTO Users (Username, Password, IsAdmin, PlainPassword) VALUES ('admin','0192023a7bbd73250516f069df18b500',true,'admin123');
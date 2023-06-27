CREATE DATABASE TGSCA;

USE TGSCA;

CREATE TABLE `Users` (
  `ID` integer PRIMARY KEY,
  `Username` varchar(255),
  `Password` varchar(255),
  `IsAdmin` bool
);

CREATE TABLE `Students` (
  `ID` integer PRIMARY KEY,
  `StudentNumber` integer,
  `UserID` integer,
  `Email` varchar(255),
  `DateOfBirth` varchar(255),
  `GradeLevel` varchar(255),
  `ContactNumber` varchar(255)
);

CREATE TABLE `Requirements` (
  `ID` integer PRIMARY KEY,
  `StudentNumber` integer,
  `UploadedFile` longtext,
  `RequirementType` varchar(255)
);

CREATE TABLE `Subjects` (
  `ID` integer PRIMARY KEY,
  `Subject` varchar(255),
  `GradeLevel` varchar(255),
  `Schedule` varchar(255),
  `Teacher` varchar(255)
);

CREATE TABLE `Enrolled` (
  `ID` integer PRIMARY KEY,
  `StudentNumber` integer,
  `SubjectID` integer
);

CREATE TABLE `Enrollment` (
  `ID` integer PRIMARY KEY,
  `StudentNumber` integer,
  `ProgressCard` longtext,
  `ProgressCardStatus` varchar(255),
  `Form137` longtext,
  `Form137Status` varchar(255),
  `GoodMoral` longtext,
  `GoodMoralStatus` varchar(255),
  `RegistrationFee` longtext,
  `RegistrationFeeStatus` varchar(255)
);

CREATE TABLE `Appointments` (
  `ID` integer PRIMARY KEY,
  `Name` varchar(255),
  `Email` varchar(255),
  `ContactNumber` varchar(255),
  `StudentNumber` integer,
  `AppointmentType` varchar(255),
  `AppointmentDescription` longtext
);

CREATE TABLE `Teachers` (
  `ID` integer PRIMARY KEY,
  `TeacherName` varchar(255)
);

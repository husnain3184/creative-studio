-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Feb 20, 2023 at 01:18 PM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 8.0.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `creative-studio`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `id` int(11) NOT NULL,
  `name` varchar(300) NOT NULL,
  `email` varchar(100) NOT NULL,
  `username` varchar(300) NOT NULL,
  `password` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`id`, `name`, `email`, `username`, `password`) VALUES
(1, 'admin', 'admin@yopmail.com', 'gtAdmin', '$2a$10$R9v2nAWDpLq4rA2agV2.GOKsl00odpDmpI67tB2/ieKdwPQ9Bvbxm'),
(2, 'sami', 'sami@yopmail.com', 'sami', '12345'),
(3, 'husnain raza', 'husnain.raza@genetechsolutions.com', 'husnain.raza', '$2a$10$olmWq/bjwvywlKdg0uNc.uV/BlZSlY8bC1Rp27vkwr7mip3P28K4i');

-- --------------------------------------------------------

--
-- Table structure for table `angels`
--

CREATE TABLE `angels` (
  `id` int(11) NOT NULL,
  `designation` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `detail` varchar(300) NOT NULL,
  `facebook` varchar(300) NOT NULL,
  `github` varchar(300) NOT NULL,
  `dropbox` varchar(300) NOT NULL,
  `twitter` varchar(300) NOT NULL,
  `image` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `angels`
--

INSERT INTO `angels` (`id`, `designation`, `name`, `detail`, `facebook`, `github`, `dropbox`, `twitter`, `image`) VALUES
(1, 'Frontend Developer', 'husnain raza', '#', '#', '#', '#', '#', '1676880487988321600.png'),
(5, 'Frontend Developer', 'husnain raza', '7wQ79r0GJ8', 'pLvjrLOE8a', 'LniVQtSj7I', 'esyu9b2UdM', 'bJVr95T5Ou', '1676891529880623200.png');

-- --------------------------------------------------------

--
-- Table structure for table `course`
--

CREATE TABLE `course` (
  `id` int(11) NOT NULL,
  `coursename` varchar(300) NOT NULL,
  `lesson` varchar(100) NOT NULL,
  `week` varchar(100) NOT NULL,
  `price` varchar(100) NOT NULL,
  `description` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `course`
--

INSERT INTO `course` (`id`, `coursename`, `lesson`, `week`, `price`, `description`) VALUES
(6, 'CSS', '10 Lesson', '1 week', '$5.00', 'shdashd asdkaslkd askdlaskd askdlaskldk');

-- --------------------------------------------------------

--
-- Table structure for table `portfolio`
--

CREATE TABLE `portfolio` (
  `id` int(6) UNSIGNED NOT NULL,
  `portfoliotitle` varchar(30) NOT NULL,
  `portfoliourl` varchar(50) DEFAULT NULL,
  `serviceimage` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `services`
--

CREATE TABLE `services` (
  `id` int(6) UNSIGNED NOT NULL,
  `servicetitle` varchar(30) NOT NULL,
  `servicdetails` varchar(50) DEFAULT NULL,
  `serviceimage` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `somefunfucts`
--

CREATE TABLE `somefunfucts` (
  `id` int(6) UNSIGNED NOT NULL,
  `title` varchar(30) NOT NULL,
  `count` varchar(50) DEFAULT NULL,
  `somefunfuctsimage` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `speakers`
--

CREATE TABLE `speakers` (
  `id` int(11) NOT NULL,
  `speakername` varchar(100) NOT NULL,
  `facebook` varchar(300) NOT NULL,
  `instagram` varchar(300) NOT NULL,
  `twiter` varchar(300) NOT NULL,
  `linkdin` varchar(300) NOT NULL,
  `image` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `speakers`
--

INSERT INTO `speakers` (`id`, `speakername`, `facebook`, `instagram`, `twiter`, `linkdin`, `image`) VALUES
(6, 'Abdul Sami', 'https://www.facebook.com/login/', 'instagram', 'https://www.facebook.com/login/', 'https://www.facebook.com/login/', '1676447272356944500.JPG');

-- --------------------------------------------------------

--
-- Table structure for table `testimonial`
--

CREATE TABLE `testimonial` (
  `id` int(6) UNSIGNED NOT NULL,
  `name` varchar(30) NOT NULL,
  `details` varchar(30) NOT NULL,
  `count` varchar(50) DEFAULT NULL,
  `image` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `angels`
--
ALTER TABLE `angels`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `course`
--
ALTER TABLE `course`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `portfolio`
--
ALTER TABLE `portfolio`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `services`
--
ALTER TABLE `services`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `somefunfucts`
--
ALTER TABLE `somefunfucts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `speakers`
--
ALTER TABLE `speakers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `testimonial`
--
ALTER TABLE `testimonial`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `angels`
--
ALTER TABLE `angels`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `portfolio`
--
ALTER TABLE `portfolio`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `services`
--
ALTER TABLE `services`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `somefunfucts`
--
ALTER TABLE `somefunfucts`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `speakers`
--
ALTER TABLE `speakers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `testimonial`
--
ALTER TABLE `testimonial`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

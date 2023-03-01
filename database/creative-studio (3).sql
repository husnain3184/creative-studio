-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 01, 2023 at 11:37 AM
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
(3, 'husnain raza', 'husnain.raza@genetechsolutions.com', 'husnain.raza', '$2a$10$olmWq/bjwvywlKdg0uNc.uV/BlZSlY8bC1Rp27vkwr7mip3P28K4i'),
(4, 'husnain raza', 'husnain.raza@genetechsolutions.com', 'husnain.raza', '$2a$10$JPVJl2hhb2ZQJElJoTTXAOULNbjI9JVtN.pdU0ZMbjBElkT7j3OM2');

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
(9, 'FOUNDER', 'Matthew Davis', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Magni quos esse tenetur illo qui, nostrum.', '#', '#', '#', '#', '1677070797449686600.jpg'),
(10, 'CEO', 'Barbara Ross', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Magni quos esse tenetur illo qui, nostrum.', '#', '#', '#', '#', '1677070832639732700.jpg'),
(11, 'Designer', 'Karen Perry', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Magni quos esse tenetur illo qui, nostrum.', '#', '#', '#', '#', '1677070867603652700.jpg'),
(12, 'App Designer', 'Ashley Diaz', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Magni quos esse tenetur illo qui, nostrum.', '#', '#', '#', '#', '1677070900706174500.jpg'),
(13, 'Developer', 'Edward Harris', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Magni quos esse tenetur illo qui, nostrum.', '#', '#', '#', '#', '1677070930630192400.jpg'),
(14, 'Photographer', 'Brian Scott', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Magni quos esse tenetur illo qui, nostrum.', '#', '#', '#', '#', '1677070974862069600.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `blogs`
--

CREATE TABLE `blogs` (
  `id` int(255) UNSIGNED NOT NULL,
  `details` varchar(300) NOT NULL,
  `image` varchar(300) DEFAULT NULL,
  `title` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `blogs`
--

INSERT INTO `blogs` (`id`, `details`, `image`, `title`) VALUES
(4, 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Repudiandae laudantium dolor nisi obcaecati, non laboriosam asperiores doloremque tempora repellendus iure!', '1677071928601172200.jpg', 'Possimus aliquam veniam'),
(5, 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Repudiandae laudantium dolor nisi obcaecati, non laboriosam asperiores doloremque tempora repellendus iure!', '1677071952982154200.jpg', 'Repudiandae laudantium'),
(6, 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Repudiandae laudantium dolor nisi obcaecati, non laboriosam asperiores doloremque tempora repellendus iure!', '1677071974161510200.jpg', 'Laboriosam asperiores'),
(8, 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Repudiandae laudantium dolor nisi obcaecati, non laboriosam asperiores doloremque tempora repellendus iure!', '1677239448852033200.jpg', 'Laboriosam asperiores');

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
  `title` varchar(30) NOT NULL,
  `url` varchar(50) DEFAULT NULL,
  `image` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `portfolio`
--

INSERT INTO `portfolio` (`id`, `title`, `url`, `image`) VALUES
(5, 'Project Title', '#', '1677070632406830000.jpg'),
(6, 'Project Title	', '#', '1677070650098221400.jpg'),
(7, 'Project Title	', '#', '1677070664253543700.jpg'),
(8, 'Project Title	', '#', '1677070677339057200.jpg'),
(9, 'Project Title	', '#', '1677070693561285600.jpg'),
(10, 'Project Title	', '#', '1677070705318164600.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `service`
--

CREATE TABLE `service` (
  `id` int(6) UNSIGNED NOT NULL,
  `title` varchar(30) NOT NULL,
  `details` varchar(300) DEFAULT NULL,
  `image` varchar(30) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `service`
--

INSERT INTO `service` (`id`, `title`, `details`, `image`) VALUES
(4, 'corporis assumenda', 'Lorem ipsum dolor sit amet, consectetur adipisicin', '1677070303182161400.PNG'),
(5, 'Debitis amet', 'Lorem ipsum dolor sit amet, consectetur adipisicin', '1677070341137747900.PNG'),
(6, 'Libero temporibus', 'Lorem ipsum dolor sit amet, consectetur adipisicin', '1677070367063733800.PNG'),
(7, 'Perspiciatis explicabo', 'Lorem ipsum dolor sit amet, consectetur adipisicin', '1677070391801273200.PNG'),
(8, 'Poluptatum', 'Lorem ipsum dolor sit amet, consectetur adipisicin', '1677070412551474600.PNG'),
(9, 'Nihil dicta', 'Lorem ipsum dolor sit amet, consectetur adipisicin', '1677070429941838600.PNG'),
(10, 'Repellendus maxime', 'Lorem ipsum dolor sit amet, consectetur adipisicin', '1677070449797222700.PNG'),
(11, 'Sint vitae', 'Lorem ipsum dolor sit amet, consectetur adipisicin', '1677070472896106500.PNG');

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
  `name` varchar(300) NOT NULL,
  `details` varchar(300) NOT NULL,
  `image` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `testimonial`
--

INSERT INTO `testimonial` (`id`, `name`, `details`, `image`) VALUES
(8, 'Richard Reb', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Eaque doloribus autem aperiam earum nostrum omnis blanditiis corporis perspiciatis adipisci nihil.', '1677071765752207800.jpg'),
(9, 'John Doe', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Eaque doloribus autem aperiam earum nostrum omnis blanditiis corporis perspiciatis adipisci nihil.', '1677071802759761100.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(255) NOT NULL,
  `name` varchar(300) NOT NULL,
  `email` varchar(300) NOT NULL,
  `username` varchar(300) NOT NULL,
  `password` varchar(300) NOT NULL,
  `roll` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `username`, `password`, `roll`) VALUES
(20, 'husnain raza', 'husnain.raza@genetechsolutions.com', 'admin123456', '$2a$10$olmWq/bjwvywlKdg0uNc.uV/BlZSlY8bC1Rp27vkwr7mip3P28K4i', 'author');

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
-- Indexes for table `blogs`
--
ALTER TABLE `blogs`
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
-- Indexes for table `service`
--
ALTER TABLE `service`
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
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `angels`
--
ALTER TABLE `angels`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `blogs`
--
ALTER TABLE `blogs`
  MODIFY `id` int(255) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `portfolio`
--
ALTER TABLE `portfolio`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `service`
--
ALTER TABLE `service`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `speakers`
--
ALTER TABLE `speakers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `testimonial`
--
ALTER TABLE `testimonial`
  MODIFY `id` int(6) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(255) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

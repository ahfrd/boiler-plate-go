-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Jun 22, 2023 at 02:34 PM
-- Server version: 5.7.34
-- PHP Version: 7.4.21

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `asia_quest`
--

-- --------------------------------------------------------

--
-- Table structure for table `books`
--

CREATE TABLE `books` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `keyword` varchar(255) NOT NULL,
  `price` varchar(255) NOT NULL,
  `stock` int(10) NOT NULL,
  `publisher` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `books`
--

INSERT INTO `books` (`id`, `title`, `description`, `category`, `keyword`, `price`, `stock`, `publisher`) VALUES
(1, 'Naruto Shipudden', 'Petualangan naruto', 'Fantasy,Action,Romance', 'Ninja,Jepang,Anime', 'Rp. 10,000,000,0,00', 5, 'Elek'),
(2, 'Naruto', 'Petualangan naruto kecil', 'Fantasy,Action,Romance', 'Ninja,Jepang,Anime', '100.000.00,00', 5, '5'),
(6, 'One Piece Eines loby', 'Petualangan Luffy', 'Fantasy,Action,Romance', 'Bajak Laut,Jepang,Anime', 'Rp. 1,234,567', 5, '5'),
(7, 'One Piece Eines loby', 'Petualangan Luffy', 'Fantasy,Action,Romance', 'Bajak Laut,Jepang,Anime', 'Rp. 1,234,567%!(EXTRA string=,00)', 5, '5'),
(8, 'One Piece Eines loby', 'Petualangan Luffy', 'Fantasy,Action,Romance', 'Bajak Laut,Jepang,Anime', 'Rp. 1,234,567,00', 5, '5');

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `username`, `password`) VALUES
(1, 'user', 'd812ccc718c4d3560f77f0680d45512f215daa22f3faa5f703fe4f2bd6970b13');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `books`
--
ALTER TABLE `books`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

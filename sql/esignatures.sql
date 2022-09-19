-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 19 Sep 2022 pada 06.59
-- Versi server: 10.4.14-MariaDB
-- Versi PHP: 7.2.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `esignatures`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `roles`
--

CREATE TABLE `roles` (
  `role_id` int(11) NOT NULL,
  `role_name` varchar(500) NOT NULL,
  `datecreated` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `roles`
--

INSERT INTO `roles` (`role_id`, `role_name`, `datecreated`) VALUES
(1, 'admin', '2022-08-13 09:08:46'),
(2, 'user', '2022-08-13 09:08:46');

-- --------------------------------------------------------

--
-- Struktur dari tabel `signatures`
--

CREATE TABLE `signatures` (
  `signature_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `signature_hash` varchar(400) NOT NULL,
  `date_created` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `transaction_id` int(11) NOT NULL,
  `address` varchar(200) NOT NULL,
  `tx_hash` varchar(500) NOT NULL,
  `nonce` varchar(200) NOT NULL,
  `description` varchar(250) NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`transaction_id`, `address`, `tx_hash`, `nonce`, `description`, `date_created`) VALUES
(1, '7196ba1565FE0bA5186E1962745353bcE0df9598', '2867927bd5fdfdcfc4d1a13f9e1f8777d830cf06639ad10ad2c12f015bceb603', '25', 'Membuat Kontrak', '2022-09-04 19:51:43'),
(2, '00f4B01c6bD6a2E19C11F833dF6063988fad9458', 'b551d95e63935eab386e18789afa55e3bf9d5a624194bdfce298f3044a0ed75a', '26', 'Membuat Kontrak', '2022-09-05 08:57:05'),
(3, 'E8384a727F4FcD41FEe36077a8104F9FdFD2e5F4', 'de94e7e71b22df5d71f24f46c041bc4473025684d7402b98eeb069eca8398697', '27', 'Membuat Kontrak', '2022-09-05 09:31:21'),
(4, '6031d77a51b581719cED34DA4A50f732342dCcFf', 'cc822141c0fd26bdee564c3566e64f3059181ebc3fd9b309b3c5d8be3bf3284a', '28', 'Membuat Kontrak', '2022-09-05 10:29:02'),
(5, '260D32662Ac0A4C4B31c63fa9C7f18E19594e2d6', '43f08c2cf156f43f545ed0f219167efc3abb5ba1768dfcb8acd28f878c6a29d3', '29', 'Membuat Kontrak', '2022-09-05 20:29:23'),
(6, '1EA2dC057501B3567815537a894c894d6f810098', '2af2393123ba1086016820aa421a8d53a726f05dd0b153b70eb40cae50bb1de0', '30', 'Membuat Kontrak', '2022-09-12 19:27:04'),
(7, '7FEa8f4d483A5C1078091B63C735Cc39d8E2e8b2', '9f679dce9a88aee0bd833268e9108b42f0fb5c58790c64b10d4597b43608cdb4', '31', 'Membuat Kontrak', '2022-09-12 19:49:01'),
(8, '24774FF4D19840AfE78815d34610C7F7D95791F4', '8564c87933c244c3d24a6a1c415172766023f6afaffbd09ea38a6f8bfae55d92', '32', 'Membuat Kontrak', '2022-09-12 19:50:51'),
(9, '5F974d65456C702803cf10DBf662B9978Fd37604', '04ff35bde311425e208f347104e049feab7675dbbe1b841eea12d6213987b902', '33', 'Membuat Kontrak', '2022-09-13 12:49:34'),
(10, '82a062abcA8F5f3F8d03dC203802D766dE053e86', 'a8cfd6d18f81d2b351e898f7cc7ed3b0c56e04552b020ba89406d85a05c35cc6', '34', 'Membuat Kontrak', '2022-09-13 16:57:54'),
(11, '3473D8F3e993717Dea03350AB425c05e10C6c715', 'a3814df930ab7314feaa673475e7b5f9749222ba82beb08ce485b7d152f315be', '35', 'Membuat Kontrak', '2022-09-13 17:20:51'),
(12, '97302f8D2dF5EBE06bb349834dc557539e8FBB0B', '19643974565e071143f6eed12212ab622e343e5b2d243e0c391b36e7d2b01a72', '36', 'Membuat Kontrak', '2022-09-13 19:49:42'),
(13, '23f999AB1e2C47Af98e7c49cA7DE56fb67fbB47b', '477761b48b2990eeec55fdf3de124468287a6fbf18bbd263c676bc6fb8ef46cb', '37', 'Membuat Kontrak', '2022-09-13 20:36:48'),
(14, '8b38ebF99C0072c04339461d17040AEB44A70D79', 'ebb6ffd4c0da7777d7bef720ea034d5a8b39f4d5f09857190849acfae9570a17', '38', 'Membuat Kontrak', '2022-09-13 20:37:41'),
(15, '1fd472825aD8D8Bb0EB2938160b4188C81777504', '47efb0eeeeb31a41d8d82e882d3ac549898a55d242523aed05f0a818e13bc26c', '39', 'Membuat Kontrak', '2022-09-13 22:04:15'),
(16, '64C56405416e08E453d6dfdd6A0296E362Cb4a39', 'caf27966c1533da25a0f01cf53772340cb72332809ddde4f21945a2bed95890f', '40', 'Membuat Kontrak', '2022-09-13 22:11:55'),
(17, 'b0610F5c6DCc622f968870d483F49Eb9b275Ca6A', '50e7ec24cf2f7fbc48281214505ee51b78f47dea793f37f149c1c97ac90d14f9', '41', 'Membuat Kontrak', '2022-09-13 22:16:10'),
(18, '6BbB89904E2dF6800a88b512de5262CD1dA2BBB2', '7df987462f2af1e68dec1c6672bb62f7bbe8c8bdbb0c8e4efabbe583fd2706ab', '42', 'Membuat Kontrak', '2022-09-14 16:07:11'),
(19, '19b3AC1E48B47ed36bCb82A20Aa21D53e83472d8', 'f6957379b692c7c36dbcbc8819aa79b219941e8e833165b0a252035d74f10680', '43', 'Membuat Kontrak', '2022-09-14 16:54:43'),
(20, 'd1B22841b17355181ee0d02bbBA5aCEfCbF8E98B', '128456c5f0ac775a20c96895c42774a615398ca3066b657406d36ffdd1088ae7', '44', 'Membuat Kontrak', '2022-09-14 17:00:46'),
(21, '2bD46F7320cCB371685E33E9649CdC8686F2d1C0', '1a69382ff62abc5cbe146deda53f078cd7a8934700b644f667b61fcd4263b8de', '45', 'Membuat Kontrak', '2022-09-14 17:10:31'),
(22, 'bE33faC9308E7377050c9613F6094bBF6d716Da0', '1c2df9caec5e0cee34d7c8a8dbbb542c95bd6ab2f3534f01c6de05324dc01a29', '46', 'Membuat Kontrak', '2022-09-14 17:49:44'),
(23, 'E82c2010106e6e7987274C6160C902E65049fDD8', '7a30e8148cdb61f192c1711fab75c51fdf39350d276e6844550c6d4c1f3e023f', '47', 'Membuat Kontrak', '2022-09-14 18:00:05'),
(24, 'B1E45dbb273699253b6cDcE99CF162cFCea4ac67', 'a15804997481cdaa51260ed8b4b970d5abd5eddee903988d4b50bfa67a9ea3fb', '48', 'Membuat Kontrak', '2022-09-14 18:01:18'),
(25, '84e93E4ED7E300C40D865EEFf052eFdf1f267b46', 'f33b6bbca69cbcb6a3408986ce9fa7a410694c9482c009b0dbd1b97dea4860a9', '49', 'Membuat Kontrak', '2022-09-14 18:07:17'),
(26, 'cb9027F41c0471830bea32D806981D0C38C77B5C', '6f8c4d48abf65f55cca835a58c0289b0f999088e5222a811685ef8152295df2d', '50', 'Membuat Kontrak', '2022-09-14 18:08:30'),
(27, '81aa9A7fc7edC4c0f56b709145557DecC92049Cb', 'f9dc952caf398a86c290a086ca05d2d52179ace56958e49aa64b766bbcc8859a', '51', 'Membuat Kontrak', '2022-09-14 18:46:32'),
(28, 'aa3EfdaC2bd3E618F05C33562952c8DA4C9E8B19', 'f9a7ef9dbe5bd10a0d3a45c6d230aae16efe65e073ceb8180fd50a7cae16af45', '52', 'Membuat Kontrak', '2022-09-14 18:48:26'),
(29, 'DB2967B7DD48104F884996F7e09A980d419439D1', '3135bea48310a6a8f9a4980ecd02e6374106188cdaa97f12b389800eef42ff3c', '53', 'Membuat Kontrak', '2022-09-14 18:49:31'),
(30, '1362c9BD7361487d14b77BD38e9DFE17925fd638', '3aefac00b2eadd1762486721b90b570bdb5d60b7eafda456dbf95e7ade8930c9', '54', 'Membuat Kontrak', '2022-09-14 18:52:46'),
(31, '0518103209fD0A1Aa8d2d18ccD4ba7934aF9F241', 'a0d6f6e7e567f6eeb04a636d6ac28c640b2c49d97de54f4f4f06cc690e587c6d', '55', 'Membuat Kontrak', '2022-09-14 18:53:36'),
(32, '2a9F63cfc846655e9A36Ca3eE0B2D2fAB387069C', 'ed186ba3ff390afe0ebcb310a7f027504d830a126327a1ee34c635be0745da12', '56', 'Membuat Kontrak', '2022-09-14 18:54:52'),
(33, 'A773D815d848dd20487E78E9914563D5025EDAd6', 'cb3ef4736a5f673841dc9b14fc1edc68be9432fed689b18b1f7d5597e0827742', '57', 'Membuat Kontrak', '2022-09-14 18:58:00'),
(34, 'b9c341Ef4344cEC16CC4f5262799Ca76e6cD257B', '673a6290257a202bbcb4b54c50b1605e9595146409bf1f80798e8b5578d97324', '58', 'Membuat Kontrak', '2022-09-14 19:21:02'),
(35, 'C455e206C31A6C9ec45d1DC4B960567BaFBe1617', '44df526354478953a39871f26800c404d8645f4bf3120df92b2c81e5c4a30270', '59', 'Membuat Kontrak', '2022-09-14 20:53:42'),
(36, '4dBb4408b9AE2Af591e82A88BD2b34CCb7C2dea2', '762c3b5dab477ab8ec7fa2faa69b2a79ee346700a475d526e337f7ae819125bc', '61', 'Membuat Kontrak', '2022-09-14 21:04:24'),
(37, '185a6b36e8b429611cb0ff4D86a828E717A4F0dD', 'd002085bacca88c4b3a4c85871f58efee55ff12a8bc61f83bab97649105db595', '62', 'Transfer Ether', '2022-09-14 21:04:55'),
(38, '141480B5Db45306A0cE3AEE8E77C090F6e88D574', '1bdc4c79c1a543d4d22bc8886350102ecec09e5837e7859ebfdf7b5532cf00b7', '63', 'Membuat Kontrak', '2022-09-19 09:14:16'),
(39, '33F9FAfC5679FfF4ff1931181c378bCeedDb918E', 'dd932bac1762d06035195b71f65a1fa3f2db8337de3687e7f5b5cb659f9c9fa6', '64', 'Membuat Kontrak', '2022-09-19 09:25:42'),
(40, '440dA5b11FC04130AfDb227b7C599a203891619F', '5877b6e0a4ea568e6a7e68fea6e2703a16408eb2a459a3cd10ee3598da149020', '65', 'Membuat Kontrak', '2022-09-19 09:28:25');

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL,
  `idsignature` varchar(500) NOT NULL,
  `name` varchar(250) NOT NULL,
  `email` varchar(250) NOT NULL,
  `phone` varchar(250) NOT NULL,
  `identity_card` varchar(400) NOT NULL,
  `password` varchar(500) NOT NULL,
  `publickey` varchar(500) NOT NULL,
  `role_id` int(11) NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`user_id`, `idsignature`, `name`, `email`, `phone`, `identity_card`, `password`, `publickey`, `role_id`, `date_created`) VALUES
(1, 'rizwijaya', 'Rizqi Wijaya', 'rizwijaya58@gmail.com', '085606394242', 'wOxefbXgtGjYaqjZTi+Lz4O9RJA4UcdYfOzKHA', '$2a$04$ye376mJqzkp9E5nWdBZDqeDkMMxigq8Ist72X.mNp2WRdV6rcwp9e', 'kwMbge2XbO7YBujq6vSDhJ/XVILQjl7k+/KY/Sdu2xEgBz8pDS4LXyWsqnvLgVtgJaUYCskucH83sioa38Ad98KK70AScA', 2, '2022-08-18 08:37:04'),
(2, 'rizqi', 'Rizwijaya', 'rizwijaya241@gmail.com', '085606394242', 'wOxefbXgtGjYaqjZTi+Lz4O9RJA4UcdYfOzKHA', '$2a$04$OfdyMolSbSCYndtt42K1butahq8H6XzGtL1PxGg/XvCXDIecBUpaW', '0zbEuc95OL6MxbhxIjQGhfNFCY0N4NcSmgnDA/OIH8F6wAvZ0hFg0DL+QmMH2uztqzkFElanW8gWdieD9Jsg6QLA8Hz8gQ', 2, '2022-08-22 11:15:24'),
(3, 'RioFerd', 'Rio Ferdinan', 'rio@gmail.com', '087628229122', 'wOxefbXgtGjYaqjZTi+Lz4O9RJA4UcdYfOzKHA', '$2a$04$88XELSJvhYwwsxxMfVgVRu5yYQ1vlrJzFKXcUPQDnsoRJUqhxc7A6', 'L2qzhxlp3SrcSRRv41k/30lV7UBvbmo3rdcNHce10Ylwfy3seRL2F9jZXnUerd7rM1gyKTaNohX5cdeGrT5yboanj+TGEg', 2, '2022-08-22 12:13:21'),
(14, 'Sutama23', 'Surya Utama', 'surya@gmail.com', '083721829233', 'ctFJDZ96V51UnXL5b6Cy8igrAYulB0Uc/jdLpg', '$2a$04$dsadQLzmfc9o5J6nI8YbTOrPFcdXmNAazR07pDCnTmImX0YxEgsFe', 'vhbpIFvc7Okent6mCOp9i3o4l3MV3be5SbPVpPijiUFw5xq00AkEBumHSbbyTYjVVDdTIRf/bT/at61esYWUhrnL2xhjcg', 2, '2022-09-14 17:01:34'),
(15, 'sembada321', 'Sembada Rusa', 'semrusa@gmail.com', '087382729121', 'H0/khs/ige1xvNREO0JAPUQk4rQCeCwD8XQ/JA', '$2a$04$MNzOhqiWntw14fAlML5NreoX.jXOSaE1YRh6LP3nTWM4p8oUVKmUi', 'Mn2OAvhmwtEtBAlFEe/AbdMvZP7YO9qCCWMuyRHKqjJvvhnt6oW7lHWtXgP9GuFrR1etSAJ8A0CHKb9sbWF8+Z1or6bSJQ', 2, '2022-09-14 17:11:26');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`role_id`);

--
-- Indeks untuk tabel `signatures`
--
ALTER TABLE `signatures`
  ADD PRIMARY KEY (`signature_id`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`transaction_id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`user_id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `roles`
--
ALTER TABLE `roles`
  MODIFY `role_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `signatures`
--
ALTER TABLE `signatures`
  MODIFY `signature_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT untuk tabel `transactions`
--
ALTER TABLE `transactions`
  MODIFY `transaction_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

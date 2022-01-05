/*
  Warnings:

  - You are about to drop the column `content` on the `profile` table. All the data in the column will be lost.
  - Added the required column `address` to the `Profile` table without a default value. This is not possible if the table is not empty.
  - Added the required column `image` to the `Profile` table without a default value. This is not possible if the table is not empty.
  - Added the required column `phone` to the `Profile` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE `profile` DROP COLUMN `content`,
    ADD COLUMN `address` VARCHAR(191) NOT NULL,
    ADD COLUMN `image` VARCHAR(191) NOT NULL,
    ADD COLUMN `phone` VARCHAR(191) NOT NULL;

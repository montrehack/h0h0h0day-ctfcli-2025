#!/bin/bash

size=50

# Resize elf
convert down.png -resize ${size}x${size}\! down_.png
convert down-hat.png -resize ${size}x${size}\! down-hat_.png
convert up.png -resize ${size}x${size}\! up_.png
convert up-hat.png -resize ${size}x${size}\! up-hat_.png
convert right.png -resize ${size}x${size}\! right_.png
convert right-hat.png -resize ${size}x${size}\! right-hat_.png

# Create left by flipping right
convert -flop right.png left.png
convert left.png -resize ${size}x${size}\! left_.png
convert -flop right-hat.png left-hat.png
convert left-hat.png -resize ${size}x${size}\! left-hat_.png

# Add snow to elf
cp up_.png up__.png
cp up-hat_.png up-hat__.png
cp down_.png down__.png
cp down-hat_.png down-hat__.png
cp left_.png left__.png
cp left-hat_.png left-hat__.png
cp right_.png right__.png
cp right-hat_.png right-hat__.png

# Add snow to trees
convert snow_.png tree-1_.png -composite tree-1__.png
convert snow_.png tree-2_.png -composite tree-2__.png
convert snow_.png tree-3_.png -composite tree-3__.png

# Resize rip
convert rip.png -resize ${size}x${size}\! rip_.png

# Add snow to rip
convert snow_.png rip_.png -composite rip__.png

# Resize wood
convert wood.png -resize ${size}x${size}\! wood__.png

# Table
convert table.png -resize ${size}x${size}\! table_.png
convert wood__.png table_.png -composite table__.png 

# Bed
convert bed.png -resize ${size}x${size}\! bed_.png
convert wood__.png bed_.png -composite bed__.png

# Chest
convert chest.png -resize ${size}x${size}\! chest_.png
convert wood__.png chest_.png -composite chest__.png

# Wall
convert wall.png -resize ${size}x${size}\! wall__.png

# Door
convert door.png -resize ${size}x${size}\! door__.png

# Hat
convert hat.png -resize ${size}x${size}\! hat_.png
convert snow_.png hat_.png -composite hat__.png

# Gift
convert gift.jpg -resize 30x30\! gift_.png
convert snow_.png gift_.png -gravity center -composite gift__.png

# Scroll
convert scroll.png -resize 30x30\! scroll_.png
convert snow_.png scroll_.png -gravity center -composite scroll__.png

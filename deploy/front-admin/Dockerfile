# Use an official Node.js runtime as a parent image
FROM node:18.15-alpine as builder

# Set the working directory to /app
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY /root/front-admin/. .

# Install dependencies
RUN npm install



# Build the React app for production
RUN npm run build

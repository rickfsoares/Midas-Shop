package com.midasshop.userservice.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.security.core.userdetails.UserDetails;

import com.midasshop.userservice.model.User;


public interface UserRepository extends JpaRepository<User, Integer>{
    UserDetails findByEmail(String email);
}

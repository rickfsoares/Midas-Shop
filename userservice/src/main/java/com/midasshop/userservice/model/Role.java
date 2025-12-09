package com.midasshop.userservice.model;

import java.util.HashSet;
import java.util.Set;

import org.springframework.security.core.GrantedAuthority;

import jakarta.persistence.Entity;
import jakarta.persistence.EnumType;
import jakarta.persistence.Enumerated;
import jakarta.persistence.Id;
import jakarta.persistence.ManyToMany;
import jakarta.persistence.Table;
import lombok.Data;
import lombok.NoArgsConstructor;

@Entity
@Data
@Table
@NoArgsConstructor
public class Role implements GrantedAuthority{

    @Id
    private int id;

    @Enumerated(EnumType.STRING)
    private RoleEnum name;

    @ManyToMany(mappedBy="roles")
    private Set<User> users= new HashSet<>(); 

    public Role(RoleEnum name){
        this.name = name;
    }

    @Override
    public String getAuthority() {
        return this.name.name();
    }

}

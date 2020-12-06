package com.github.matheushr97;

import io.quarkus.hibernate.orm.panache.PanacheEntity;

import javax.persistence.Column;
import javax.persistence.Entity;

@Entity(name = "heroes")
public class Hero extends PanacheEntity {

    private String name;

    @Column(name = "special_powers")
    private String specialPowers;

    private String publisher;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getSpecialPowers() {
        return specialPowers;
    }

    public void setSpecialPowers(String specialPowers) {
        this.specialPowers = specialPowers;
    }

    public String getPublisher() {
        return publisher;
    }

    public void setPublisher(String publisher) {
        this.publisher = publisher;
    }
}

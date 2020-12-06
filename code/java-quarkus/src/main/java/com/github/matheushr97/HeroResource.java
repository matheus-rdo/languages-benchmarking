package com.github.matheushr97;

import javax.ws.rs.Consumes;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import java.util.List;

@Path("/heroes")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
public class HeroResource {

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public List<Hero> hello() {
        List<Hero> heroes = Hero.listAll();
        return heroes;
    }
}
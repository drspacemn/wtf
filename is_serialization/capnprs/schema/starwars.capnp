@0xc1da187e3c0d97cd;

interface StarWars {
    struct Human {
      id @0 :Text;
      name @1 :Text;
      homePlanet @2 :Text;
      appearsIn @3 :AppearsIn;

      enum AppearsIn {
        newHope @0;
        empire @1;
        jedi @2;
      }
    }

    showHuman @0 (id: Text) -> Human;
    createHuman @1 Human -> Human;
}
<template>
  <Minimal>
    <div class="row">
      <div class="col-12">
        <div class="card">
          <div class="card-body">
            <form @submit="submit">
              <div class="form-group">
                <label for="email">Email de connexion</label>
                <input
                  type="email"
                  class="form-control"
                  id="email"
                  aria-describedby="emailHelp"
                  placeholder="Entrer votre email"
                  required
                  v-model="email"
                />
                <small id="emailHelp" class="form-text text-muted">Identifiant de connexion Awise</small>
              </div>
              <div class="form-group">
                <label for="password">Mot de passe</label>
                <input
                  type="password"
                  class="form-control"
                  id="password"
                  placeholder="Entrer votre mot de passe"
                  v-model="password"
                />
              </div>
              <div class="form-group form-check">
                <input type="checkbox" class="form-check-input" id="remember" v-model="remember" />
                <label class="form-check-label" for="remember">Se souvenir de moi</label>
              </div>
              <button type="submit" class="btn btn-primary">Se connecter</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </Minimal>
</template>

<script>
import Minimal from "../../layout/minimal";
import Socket from "../../plugings/socket";
import { fetch } from "../../plugings/request";
export default {
  name: "Home",
  components: {
    Minimal
  },
  data: function() {
    return { email: "", password: "", remember: true };
  },
  methods: {
    submit: function(e) {
      e.preventDefault();
    }
  }
};
</script>

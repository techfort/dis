<template>
  <div>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark fixed-top">
      <a class="navbar-brand" href="#">DIS</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        <ul class="navbar-nav mr-auto">
          <li class="nav-item">
            <a class="nav-link"><router-link to="/watch">Watch</router-link></a>
          </li>
          <li class="nav-item">
            <a class="nav-link"><router-link to="/query">Query</router-link></a>
          </li>
          <li class="nav-item">
            <a class="nav-link"><router-link to="/insert">Insert</router-link></a>
          </li>
          <li class="nav-item">
            <a class="nav-link"><router-link to="/sets">Sets</router-link></a>
          </li>
        </ul>
      </div>
    </nav>
    <div class="container-fluid" id="mainview">
      <router-view></router-view>
    </div>
  </div>
</template>

<style>
@import url('https://fonts.googleapis.com/css?family=Inconsolata');
html, body {
  font-family: 'Inconsolata', monospace !important;
  color: #333;
  letter-spacing: 2px !important;
}
body {
  padding-top: 5rem;
}
.navbar-dark {
  background: #6D0010 !important;
}

.nav-link{
  color: white;
  font-size: 1.1em;
  font-weight: 800;
}

.nav-link a:visited {
  color: white !important;
}

* {
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
}

.pure-menu {
  border-bottom: 1px solid #333;
}

.pure-menu-item {
  vertical-align: middle;
  padding: 0 2em 0;
}
.insertform {
    padding: 1rem;
    border: 1px solid #ccc;
    margin: 1rem;
}
.insertform input, .insertform button {
    margin: 0.5rem;
    padding: 0.5rem;
}

.badge {
  font-variant: small-caps;
}
h6 {
  font-weight: 600;
}
</style>

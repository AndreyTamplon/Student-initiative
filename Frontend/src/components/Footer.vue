<template>
  <footer class="footer_area section_padding_130_0">
    <div class="container">
      <div class="row">
        <div class="col-12 col-sm-6 col-lg-4">
          <div class="single-footer-widget section_padding_0_130">
            <div class="footer-logo mb-3"></div>
            <p>The bridge between the student and the dean's office.</p>
            <br>
            <div class="copywrite-text mb-5">
              <p class="mb-0">Made with love by <a class="ml-1" href="https://github.com/AndreyTamplon/">Andrey
                Tamplon</a></p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Footer"
}
</script>

<style scoped>

.footer_area {
  position: relative;
  z-index: 1;
  overflow: hidden;
  webkit-box-shadow: 0 8px 48px 8px rgba(47, 91, 234, 0.175);
  box-shadow: 0 4px 24px 4px rgba(47, 91, 234, 0.175);
  padding: 10px;
  background-color: #ffffff;
}

.footer_area .row {
  margin-left: 38%;
  margin-right: -25px;
}

.footer_area .row .col-12,
.footer_area .row .col-lg-4,
.footer_area .row .col-sm-6,


.single-footer-widget .copywrite-text a {
  color: #747794;
  font-size: 1rem;
}

.single-footer-widget .copywrite-text a:hover,
.single-footer-widget .copywrite-text a:focus {
  color: #3f43fd;
}

.single-footer-widget .footer_menu li a {
  color: #747794;
  display: block;
  font-size: 1rem;
}

.single-footer-widget .footer_menu li a:hover,
.single-footer-widget .footer_menu li a:focus {
  color: #3f43fd;
}

.single-footer-widget .footer_menu li:last-child a {
}
</style>
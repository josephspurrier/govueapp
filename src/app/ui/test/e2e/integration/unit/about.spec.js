/* eslint-disable no-undef */
import About from '@/pages/about.vue'

const mount = require('cypress-vue-unit-test')

describe('about.vue', function() {
  beforeEach(mount(About))

  it('renders text', function() {
    cy.wrap(Cypress.vue).should('not.be.undefined')
  })

  it('has a few headings', function() {
    cy.get('h1').contains('About')
    cy.get('h2').contains(
      'This shows you how to build a website using Nuxt, Vue, and Bulma.'
    )
  })
})

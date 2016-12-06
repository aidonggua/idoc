import { IdocPage } from './app.po';

describe('idoc App', function() {
  let page: IdocPage;

  beforeEach(() => {
    page = new IdocPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});

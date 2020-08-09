from bs4 import BeautifulSoup
import requests
import re
import pandas as pd

final_price_url = "https://www.hemnet.se/salda/bostader?housing_form_groups=apartments&location_ids=17744"


class ReadSoup:
    def __init__(self, expand_hits):
        self.test_clean = []
        self.expand_hits = expand_hits
        self.parse(expand_hits)

    def parse(self, expand_hits):
        commune_name = \
            expand_hits.find("div", {"class": "sold-property-listing__location"}) \
                .findAll("span",
                         {
                             "class": "item-link"})
        if len(commune_name) == 2:
            place_name = \
                expand_hits.find("div", {"class": "sold-property-listing__location"}).findAll("span",
                                                                                              {"class": "item-link"})[
                    1]

            size = expand_hits.find('div', class_='sold-property-listing__size')
            monthly = expand_hits.find('div', class_='sold-property-listing__price')
            final_str = place_name.getText() + ", " + monthly.getText() + ", " + size.getText()
            test_unclean = final_str.splitlines()
            test_unclean = [e.strip() for e in test_unclean]
            self.test_clean = [e for e in test_unclean if len(e) != 0]
        else:
            print(commune_name)
            Exception("Failed to get commune name")

    def get_region(self):
        return self.test_clean[0].replace(",", "")

    def get_price(self):
        return int(re.sub('[^0-9]', '', self.test_clean[2]))

    def get_price_per_size(self):
        return int(re.sub('[^0-9]', '', self.test_clean[4]))

    def get_size(self):
        return float(self.test_clean[6].replace(u'\xa0', ' ').replace(" m²", "").replace(",", "."))

    def get_room(self):
        return float(self.test_clean[7].replace(u'\xa0', ' ').replace(" rum", "").replace(",", "."))

    def get_rent(self):
        return int(re.sub('[^0-9]', '', self.test_clean[8]))


def scrape(price, size):
    clean_apartments = []
    missed_clean_up = []
    for i in range(1, 51):
        r = requests.get(
            final_price_url + "&selling_price_min=" + price[0] + "&selling_price_max=" + price[
                1] + "living_area_min=" + size[0] + "&living_area_max=" + size[1] + "&page=" + str(
                i))
        if r.status_code != 200:
            print(r.content)
            continue
        data = r.content
        soup = BeautifulSoup(data, features="html.parser")
        required_elements = soup.findAll("div", {"class": "sold-property-listing"})
        for e in required_elements:
            un_marshaled = ReadSoup(e)
            if len(un_marshaled.test_clean) == 9:
                try:
                    clean_apartments.append(
                        {'region': un_marshaled.get_region(),
                         'price': un_marshaled.get_price(),
                         'price_per_size': un_marshaled.get_price_per_size(),
                         'size': un_marshaled.get_size(),
                         'rooms': un_marshaled.get_room(),
                         'rent': un_marshaled.get_rent()
                         })
                except (ValueError, Exception) as e:
                    print("Oops!  That was no valid number.  Try again...")
                    print(e)
                    missed_clean_up.append(un_marshaled.test_clean)
            else:
                missed_clean_up.append(un_marshaled.test_clean)

    append_to_csv(clean_apartments, price[1] + "" + size[1])
    print(len(missed_clean_up))
    return clean_apartments


def append_to_csv(raw_data, price_max):
    df = pd.DataFrame(data=raw_data)
    df.to_csv('test_data_' + str(price_max) + '.csv', encoding='utf-8', index=False)
    pass


if __name__ == '__main__':
    full_clean_apartments = []
    price_ranges = [["0", "2500000"], ["2500001", "5000000"], ["5000001", "10000000"]]
    space_range = [["0", "30"], ["31", "50"], ["51", "70"], ["71", "100"]]
    for bracket in price_ranges:
        for space in space_range:
            full_clean_apartments.extend(scrape(bracket, space))

    append_to_csv(full_clean_apartments, "full")

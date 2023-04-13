import datetime
import time
from phue import Bridge

# Hier importeren we de benodigde modules. We hebben datetime nodig om de tijd te bepalen en time om te wachten tussen de tijden door. De phue module is nodig om met de Philips Hue Bridge te verbinden.

# Configuratie van de Philips Hue Bridge
bridge_ip = "192.168.1.2" # Het IP-adres van de Philips Hue Bridge
username = "DitIsMijnUsername" # De gebruikersnaam waarmee we verbinding maken met de Philips Hue Bridge
b = Bridge(bridge_ip, username) # Maak een verbinding met de Philips Hue Bridge

# Instellen van de lichtsferen op basis van tijd
light_schemes = {
    "morning": {"light_id": 1, "brightness": 100, "hue": 46920, "saturation": 254},
    "evening": {"light_id": 1, "brightness": 100, "hue": 14910, "saturation": 254}
}

# Hier definiëren we een dictionary (light_schemes) die aangeeft welke lichtsfeer bij welke tijd hoort. We hebben hier twee tijden gedefinieerd, "morning" en "evening", en bij elke tijd aangegeven welk licht we willen laten branden en met welke kleur.

# Functie om de lichten te bedienen
def set_light(light_id, brightness, hue, saturation):
    """
    Zet de verlichting aan met de opgegeven instellingen.
    """
    b.set_light(light_id, {
        "on": True,
        "bri": brightness,
        "hue": hue,
        "sat": saturation
    })

# Hier definiëren we een functie (set_light) die de lichten bedient met de opgegeven instellingen. Deze functie zal later worden gebruikt om de lichten in de juiste lichtsfeer te zetten.

# Main
while True:
    current_time = datetime.datetime.now().time() # Bepaal de huidige tijd
    if current_time >= datetime.time(6, 0) and current_time < datetime.time(8, 0): # Als het tussen 6:00 en 8:00 uur is, zet dan de lichten op "morning"
        set_light(**light_schemes["morning"])
    elif current_time >= datetime.time(21, 0) or current_time < datetime.time(6, 0)): # Als het na 21:00 uur is of voor 6:00 uur, zet dan de lichten op "evening"
        set_light(**light_schemes["evening"])
    else:
        b.set_light(1, "on", False) # Anders zet het licht uit

    time.sleep(60) # Wacht 60 seconden voordat we de tijd opnieuw controleren

package models

import (
	"fmt"
	"net/url"
	"strings"
)

// Defines values for ForecastParamsHourly.
const (
	ForecastParamsHourlyAlbedo                           ForecastParamsHourly = "albedo"
	ForecastParamsHourlyApparentTemperature              ForecastParamsHourly = "apparent_temperature"
	ForecastParamsHourlyBoundaryLayerHeight              ForecastParamsHourly = "boundary_layer_height"
	ForecastParamsHourlyCape                             ForecastParamsHourly = "cape"
	ForecastParamsHourlyCloudCover                       ForecastParamsHourly = "cloud_cover"
	ForecastParamsHourlyCloudCover1000hPa                ForecastParamsHourly = "cloud_cover_1000hPa"
	ForecastParamsHourlyCloudCover100hPa                 ForecastParamsHourly = "cloud_cover_100hPa"
	ForecastParamsHourlyCloudCover10hPa                  ForecastParamsHourly = "cloud_cover_10hPa"
	ForecastParamsHourlyCloudCover125hPa                 ForecastParamsHourly = "cloud_cover_125hPa"
	ForecastParamsHourlyCloudCover150hPa                 ForecastParamsHourly = "cloud_cover_150hPa"
	ForecastParamsHourlyCloudCover175hPa                 ForecastParamsHourly = "cloud_cover_175hPa"
	ForecastParamsHourlyCloudCover200hPa                 ForecastParamsHourly = "cloud_cover_200hPa"
	ForecastParamsHourlyCloudCover225hPa                 ForecastParamsHourly = "cloud_cover_225hPa"
	ForecastParamsHourlyCloudCover250hPa                 ForecastParamsHourly = "cloud_cover_250hPa"
	ForecastParamsHourlyCloudCover275hPa                 ForecastParamsHourly = "cloud_cover_275hPa"
	ForecastParamsHourlyCloudCover300hPa                 ForecastParamsHourly = "cloud_cover_300hPa"
	ForecastParamsHourlyCloudCover350hPa                 ForecastParamsHourly = "cloud_cover_350hPa"
	ForecastParamsHourlyCloudCover400hPa                 ForecastParamsHourly = "cloud_cover_400hPa"
	ForecastParamsHourlyCloudCover450hPa                 ForecastParamsHourly = "cloud_cover_450hPa"
	ForecastParamsHourlyCloudCover500hPa                 ForecastParamsHourly = "cloud_cover_500hPa"
	ForecastParamsHourlyCloudCover50hPa                  ForecastParamsHourly = "cloud_cover_50hPa"
	ForecastParamsHourlyCloudCover550hPa                 ForecastParamsHourly = "cloud_cover_550hPa"
	ForecastParamsHourlyCloudCover600hPa                 ForecastParamsHourly = "cloud_cover_600hPa"
	ForecastParamsHourlyCloudCover650hPa                 ForecastParamsHourly = "cloud_cover_650hPa"
	ForecastParamsHourlyCloudCover700hPa                 ForecastParamsHourly = "cloud_cover_700hPa"
	ForecastParamsHourlyCloudCover750hPa                 ForecastParamsHourly = "cloud_cover_750hPa"
	ForecastParamsHourlyCloudCover800hPa                 ForecastParamsHourly = "cloud_cover_800hPa"
	ForecastParamsHourlyCloudCover850hPa                 ForecastParamsHourly = "cloud_cover_850hPa"
	ForecastParamsHourlyCloudCover900hPa                 ForecastParamsHourly = "cloud_cover_900hPa"
	ForecastParamsHourlyCloudCover925hPa                 ForecastParamsHourly = "cloud_cover_925hPa"
	ForecastParamsHourlyCloudCover950hPa                 ForecastParamsHourly = "cloud_cover_950hPa"
	ForecastParamsHourlyCloudCoverHigh                   ForecastParamsHourly = "cloud_cover_high"
	ForecastParamsHourlyCloudCoverLow                    ForecastParamsHourly = "cloud_cover_low"
	ForecastParamsHourlyCloudCoverMid                    ForecastParamsHourly = "cloud_cover_mid"
	ForecastParamsHourlyConvectiveCloudBase              ForecastParamsHourly = "convective_cloud_base"
	ForecastParamsHourlyConvectiveCloudTop               ForecastParamsHourly = "convective_cloud_top"
	ForecastParamsHourlyConvectiveInhibition             ForecastParamsHourly = "convective_inhibition"
	ForecastParamsHourlyDewPoint1000hPa                  ForecastParamsHourly = "dew_point_1000hPa"
	ForecastParamsHourlyDewPoint100hPa                   ForecastParamsHourly = "dew_point_100hPa"
	ForecastParamsHourlyDewPoint10hPa                    ForecastParamsHourly = "dew_point_10hPa"
	ForecastParamsHourlyDewPoint125hPa                   ForecastParamsHourly = "dew_point_125hPa"
	ForecastParamsHourlyDewPoint150hPa                   ForecastParamsHourly = "dew_point_150hPa"
	ForecastParamsHourlyDewPoint175hPa                   ForecastParamsHourly = "dew_point_175hPa"
	ForecastParamsHourlyDewPoint200hPa                   ForecastParamsHourly = "dew_point_200hPa"
	ForecastParamsHourlyDewPoint225hPa                   ForecastParamsHourly = "dew_point_225hPa"
	ForecastParamsHourlyDewPoint250hPa                   ForecastParamsHourly = "dew_point_250hPa"
	ForecastParamsHourlyDewPoint275hPa                   ForecastParamsHourly = "dew_point_275hPa"
	ForecastParamsHourlyDewPoint2m                       ForecastParamsHourly = "dew_point_2m"
	ForecastParamsHourlyDewPoint300hPa                   ForecastParamsHourly = "dew_point_300hPa"
	ForecastParamsHourlyDewPoint350hPa                   ForecastParamsHourly = "dew_point_350hPa"
	ForecastParamsHourlyDewPoint400hPa                   ForecastParamsHourly = "dew_point_400hPa"
	ForecastParamsHourlyDewPoint450hPa                   ForecastParamsHourly = "dew_point_450hPa"
	ForecastParamsHourlyDewPoint500hPa                   ForecastParamsHourly = "dew_point_500hPa"
	ForecastParamsHourlyDewPoint50hPa                    ForecastParamsHourly = "dew_point_50hPa"
	ForecastParamsHourlyDewPoint550hPa                   ForecastParamsHourly = "dew_point_550hPa"
	ForecastParamsHourlyDewPoint600hPa                   ForecastParamsHourly = "dew_point_600hPa"
	ForecastParamsHourlyDewPoint650hPa                   ForecastParamsHourly = "dew_point_650hPa"
	ForecastParamsHourlyDewPoint700hPa                   ForecastParamsHourly = "dew_point_700hPa"
	ForecastParamsHourlyDewPoint750hPa                   ForecastParamsHourly = "dew_point_750hPa"
	ForecastParamsHourlyDewPoint800hPa                   ForecastParamsHourly = "dew_point_800hPa"
	ForecastParamsHourlyDewPoint850hPa                   ForecastParamsHourly = "dew_point_850hPa"
	ForecastParamsHourlyDewPoint900hPa                   ForecastParamsHourly = "dew_point_900hPa"
	ForecastParamsHourlyDewPoint925hPa                   ForecastParamsHourly = "dew_point_925hPa"
	ForecastParamsHourlyDewPoint950hPa                   ForecastParamsHourly = "dew_point_950hPa"
	ForecastParamsHourlyDiffuseRadiation                 ForecastParamsHourly = "diffuse_radiation"
	ForecastParamsHourlyDiffuseRadiationInstant          ForecastParamsHourly = "diffuse_radiation_instant"
	ForecastParamsHourlyDirectNormalIrradiance           ForecastParamsHourly = "direct_normal_irradiance"
	ForecastParamsHourlyDirectNormalIrradianceInstant    ForecastParamsHourly = "direct_normal_irradiance_instant"
	ForecastParamsHourlyDirectRadiation                  ForecastParamsHourly = "direct_radiation"
	ForecastParamsHourlyDirectRadiationInstant           ForecastParamsHourly = "direct_radiation_instant"
	ForecastParamsHourlyEt0FaoEvapotranspiration         ForecastParamsHourly = "et0_fao_evapotranspiration"
	ForecastParamsHourlyEvapotranspiration               ForecastParamsHourly = "evapotranspiration"
	ForecastParamsHourlyFreezingLevelHeight              ForecastParamsHourly = "freezing_level_height"
	ForecastParamsHourlyFreezingRainProbability          ForecastParamsHourly = "freezing_rain_probability"
	ForecastParamsHourlyGeopotentialHeight1000hPa        ForecastParamsHourly = "geopotential_height_1000hPa"
	ForecastParamsHourlyGeopotentialHeight100hPa         ForecastParamsHourly = "geopotential_height_100hPa"
	ForecastParamsHourlyGeopotentialHeight10hPa          ForecastParamsHourly = "geopotential_height_10hPa"
	ForecastParamsHourlyGeopotentialHeight125hPa         ForecastParamsHourly = "geopotential_height_125hPa"
	ForecastParamsHourlyGeopotentialHeight150hPa         ForecastParamsHourly = "geopotential_height_150hPa"
	ForecastParamsHourlyGeopotentialHeight175hPa         ForecastParamsHourly = "geopotential_height_175hPa"
	ForecastParamsHourlyGeopotentialHeight200hPa         ForecastParamsHourly = "geopotential_height_200hPa"
	ForecastParamsHourlyGeopotentialHeight225hPa         ForecastParamsHourly = "geopotential_height_225hPa"
	ForecastParamsHourlyGeopotentialHeight250hPa         ForecastParamsHourly = "geopotential_height_250hPa"
	ForecastParamsHourlyGeopotentialHeight275hPa         ForecastParamsHourly = "geopotential_height_275hPa"
	ForecastParamsHourlyGeopotentialHeight300hPa         ForecastParamsHourly = "geopotential_height_300hPa"
	ForecastParamsHourlyGeopotentialHeight350hPa         ForecastParamsHourly = "geopotential_height_350hPa"
	ForecastParamsHourlyGeopotentialHeight400hPa         ForecastParamsHourly = "geopotential_height_400hPa"
	ForecastParamsHourlyGeopotentialHeight450hPa         ForecastParamsHourly = "geopotential_height_450hPa"
	ForecastParamsHourlyGeopotentialHeight500hPa         ForecastParamsHourly = "geopotential_height_500hPa"
	ForecastParamsHourlyGeopotentialHeight50hPa          ForecastParamsHourly = "geopotential_height_50hPa"
	ForecastParamsHourlyGeopotentialHeight550hPa         ForecastParamsHourly = "geopotential_height_550hPa"
	ForecastParamsHourlyGeopotentialHeight600hPa         ForecastParamsHourly = "geopotential_height_600hPa"
	ForecastParamsHourlyGeopotentialHeight650hPa         ForecastParamsHourly = "geopotential_height_650hPa"
	ForecastParamsHourlyGeopotentialHeight700hPa         ForecastParamsHourly = "geopotential_height_700hPa"
	ForecastParamsHourlyGeopotentialHeight750hPa         ForecastParamsHourly = "geopotential_height_750hPa"
	ForecastParamsHourlyGeopotentialHeight800hPa         ForecastParamsHourly = "geopotential_height_800hPa"
	ForecastParamsHourlyGeopotentialHeight850hPa         ForecastParamsHourly = "geopotential_height_850hPa"
	ForecastParamsHourlyGeopotentialHeight900hPa         ForecastParamsHourly = "geopotential_height_900hPa"
	ForecastParamsHourlyGeopotentialHeight925hPa         ForecastParamsHourly = "geopotential_height_925hPa"
	ForecastParamsHourlyGeopotentialHeight950hPa         ForecastParamsHourly = "geopotential_height_950hPa"
	ForecastParamsHourlyGlobalTiltedIrradiance           ForecastParamsHourly = "global_tilted_irradiance"
	ForecastParamsHourlyGlobalTiltedIrradianceInstant    ForecastParamsHourly = "global_tilted_irradiance_instant"
	ForecastParamsHourlyIcePelletsProbability            ForecastParamsHourly = "ice_pellets_probability"
	ForecastParamsHourlyIsDay                            ForecastParamsHourly = "is_day"
	ForecastParamsHourlyKIndex                           ForecastParamsHourly = "k_index"
	ForecastParamsHourlyLiftedIndex                      ForecastParamsHourly = "lifted_index"
	ForecastParamsHourlyLightningDensity                 ForecastParamsHourly = "lightning_density"
	ForecastParamsHourlyLightningPotential               ForecastParamsHourly = "lightning_potential"
	ForecastParamsHourlyMassDensity8m                    ForecastParamsHourly = "mass_density_8m"
	ForecastParamsHourlyOceanCurrentDirection            ForecastParamsHourly = "ocean_current_direction"
	ForecastParamsHourlyOceanCurrentVelocity             ForecastParamsHourly = "ocean_current_velocity"
	ForecastParamsHourlyPrecipitation                    ForecastParamsHourly = "precipitation"
	ForecastParamsHourlyPrecipitationProbability         ForecastParamsHourly = "precipitation_probability"
	ForecastParamsHourlyPrecipitationType                ForecastParamsHourly = "precipitation_type"
	ForecastParamsHourlyPressureMsl                      ForecastParamsHourly = "pressure_msl"
	ForecastParamsHourlyRain                             ForecastParamsHourly = "rain"
	ForecastParamsHourlyRainProbability                  ForecastParamsHourly = "rain_probability"
	ForecastParamsHourlyRelativeHumidity1000hPa          ForecastParamsHourly = "relative_humidity_1000hPa"
	ForecastParamsHourlyRelativeHumidity100hPa           ForecastParamsHourly = "relative_humidity_100hPa"
	ForecastParamsHourlyRelativeHumidity10hPa            ForecastParamsHourly = "relative_humidity_10hPa"
	ForecastParamsHourlyRelativeHumidity125hPa           ForecastParamsHourly = "relative_humidity_125hPa"
	ForecastParamsHourlyRelativeHumidity150hPa           ForecastParamsHourly = "relative_humidity_150hPa"
	ForecastParamsHourlyRelativeHumidity175hPa           ForecastParamsHourly = "relative_humidity_175hPa"
	ForecastParamsHourlyRelativeHumidity200hPa           ForecastParamsHourly = "relative_humidity_200hPa"
	ForecastParamsHourlyRelativeHumidity225hPa           ForecastParamsHourly = "relative_humidity_225hPa"
	ForecastParamsHourlyRelativeHumidity250hPa           ForecastParamsHourly = "relative_humidity_250hPa"
	ForecastParamsHourlyRelativeHumidity275hPa           ForecastParamsHourly = "relative_humidity_275hPa"
	ForecastParamsHourlyRelativeHumidity2m               ForecastParamsHourly = "relative_humidity_2m"
	ForecastParamsHourlyRelativeHumidity300hPa           ForecastParamsHourly = "relative_humidity_300hPa"
	ForecastParamsHourlyRelativeHumidity350hPa           ForecastParamsHourly = "relative_humidity_350hPa"
	ForecastParamsHourlyRelativeHumidity400hPa           ForecastParamsHourly = "relative_humidity_400hPa"
	ForecastParamsHourlyRelativeHumidity450hPa           ForecastParamsHourly = "relative_humidity_450hPa"
	ForecastParamsHourlyRelativeHumidity500hPa           ForecastParamsHourly = "relative_humidity_500hPa"
	ForecastParamsHourlyRelativeHumidity50hPa            ForecastParamsHourly = "relative_humidity_50hPa"
	ForecastParamsHourlyRelativeHumidity550hPa           ForecastParamsHourly = "relative_humidity_550hPa"
	ForecastParamsHourlyRelativeHumidity600hPa           ForecastParamsHourly = "relative_humidity_600hPa"
	ForecastParamsHourlyRelativeHumidity650hPa           ForecastParamsHourly = "relative_humidity_650hPa"
	ForecastParamsHourlyRelativeHumidity700hPa           ForecastParamsHourly = "relative_humidity_700hPa"
	ForecastParamsHourlyRelativeHumidity750hPa           ForecastParamsHourly = "relative_humidity_750hPa"
	ForecastParamsHourlyRelativeHumidity800hPa           ForecastParamsHourly = "relative_humidity_800hPa"
	ForecastParamsHourlyRelativeHumidity850hPa           ForecastParamsHourly = "relative_humidity_850hPa"
	ForecastParamsHourlyRelativeHumidity900hPa           ForecastParamsHourly = "relative_humidity_900hPa"
	ForecastParamsHourlyRelativeHumidity925hPa           ForecastParamsHourly = "relative_humidity_925hPa"
	ForecastParamsHourlyRelativeHumidity950hPa           ForecastParamsHourly = "relative_humidity_950hPa"
	ForecastParamsHourlyRoughnessLength                  ForecastParamsHourly = "roughness_length"
	ForecastParamsHourlyRunoff                           ForecastParamsHourly = "runoff"
	ForecastParamsHourlySeaIceThickness                  ForecastParamsHourly = "sea_ice_thickness"
	ForecastParamsHourlySeaLevelHeightMsl                ForecastParamsHourly = "sea_level_height_msl"
	ForecastParamsHourlySeaSurfaceTemperature            ForecastParamsHourly = "sea_surface_temperature"
	ForecastParamsHourlyShortwaveRadiation               ForecastParamsHourly = "shortwave_radiation"
	ForecastParamsHourlyShortwaveRadiationInstant        ForecastParamsHourly = "shortwave_radiation_instant"
	ForecastParamsHourlyShowers                          ForecastParamsHourly = "showers"
	ForecastParamsHourlySnowDepth                        ForecastParamsHourly = "snow_depth"
	ForecastParamsHourlySnowDepthWaterEquivalent         ForecastParamsHourly = "snow_depth_water_equivalent"
	ForecastParamsHourlySnowHeight                       ForecastParamsHourly = "snow_height"
	ForecastParamsHourlySnowfall                         ForecastParamsHourly = "snowfall"
	ForecastParamsHourlySnowfallHeight                   ForecastParamsHourly = "snowfall_height"
	ForecastParamsHourlySnowfallProbability              ForecastParamsHourly = "snowfall_probability"
	ForecastParamsHourlySnowfallWaterEquivalent          ForecastParamsHourly = "snowfall_water_equivalent"
	ForecastParamsHourlySoilMoisture0To10cm              ForecastParamsHourly = "soil_moisture_0_to_10cm"
	ForecastParamsHourlySoilMoisture0To1cm               ForecastParamsHourly = "soil_moisture_0_to_1cm"
	ForecastParamsHourlySoilMoisture0To7cm               ForecastParamsHourly = "soil_moisture_0_to_7cm"
	ForecastParamsHourlySoilMoisture100To200cm           ForecastParamsHourly = "soil_moisture_100_to_200cm"
	ForecastParamsHourlySoilMoisture100To255cm           ForecastParamsHourly = "soil_moisture_100_to_255cm"
	ForecastParamsHourlySoilMoisture100To300cm           ForecastParamsHourly = "soil_moisture_100_to_300cm"
	ForecastParamsHourlySoilMoisture10To35cm             ForecastParamsHourly = "soil_moisture_10_to_35cm"
	ForecastParamsHourlySoilMoisture10To40cm             ForecastParamsHourly = "soil_moisture_10_to_40cm"
	ForecastParamsHourlySoilMoisture1To3cm               ForecastParamsHourly = "soil_moisture_1_to_3cm"
	ForecastParamsHourlySoilMoisture243To729cm           ForecastParamsHourly = "soil_moisture_243_to_729cm"
	ForecastParamsHourlySoilMoisture27To81cm             ForecastParamsHourly = "soil_moisture_27_to_81cm"
	ForecastParamsHourlySoilMoisture28To100cm            ForecastParamsHourly = "soil_moisture_28_to_100cm"
	ForecastParamsHourlySoilMoisture35To100cm            ForecastParamsHourly = "soil_moisture_35_to_100cm"
	ForecastParamsHourlySoilMoisture3To9cm               ForecastParamsHourly = "soil_moisture_3_to_9cm"
	ForecastParamsHourlySoilMoisture40To100cm            ForecastParamsHourly = "soil_moisture_40_to_100cm"
	ForecastParamsHourlySoilMoisture729To2187cm          ForecastParamsHourly = "soil_moisture_729_to_2187cm"
	ForecastParamsHourlySoilMoisture7To28cm              ForecastParamsHourly = "soil_moisture_7_to_28cm"
	ForecastParamsHourlySoilMoisture81To243cm            ForecastParamsHourly = "soil_moisture_81_to_243cm"
	ForecastParamsHourlySoilMoisture9To27cm              ForecastParamsHourly = "soil_moisture_9_to_27cm"
	ForecastParamsHourlySoilTemperature0To10cm           ForecastParamsHourly = "soil_temperature_0_to_10cm"
	ForecastParamsHourlySoilTemperature0To7cm            ForecastParamsHourly = "soil_temperature_0_to_7cm"
	ForecastParamsHourlySoilTemperature0cm               ForecastParamsHourly = "soil_temperature_0cm"
	ForecastParamsHourlySoilTemperature100To200cm        ForecastParamsHourly = "soil_temperature_100_to_200cm"
	ForecastParamsHourlySoilTemperature100To255cm        ForecastParamsHourly = "soil_temperature_100_to_255cm"
	ForecastParamsHourlySoilTemperature100To300cm        ForecastParamsHourly = "soil_temperature_100_to_300cm"
	ForecastParamsHourlySoilTemperature10To35cm          ForecastParamsHourly = "soil_temperature_10_to_35cm"
	ForecastParamsHourlySoilTemperature10To40cm          ForecastParamsHourly = "soil_temperature_10_to_40cm"
	ForecastParamsHourlySoilTemperature1458cm            ForecastParamsHourly = "soil_temperature_1458cm"
	ForecastParamsHourlySoilTemperature162cm             ForecastParamsHourly = "soil_temperature_162cm"
	ForecastParamsHourlySoilTemperature18cm              ForecastParamsHourly = "soil_temperature_18cm"
	ForecastParamsHourlySoilTemperature28To100cm         ForecastParamsHourly = "soil_temperature_28_to_100cm"
	ForecastParamsHourlySoilTemperature35To100cm         ForecastParamsHourly = "soil_temperature_35_to_100cm"
	ForecastParamsHourlySoilTemperature40To100cm         ForecastParamsHourly = "soil_temperature_40_to_100cm"
	ForecastParamsHourlySoilTemperature486cm             ForecastParamsHourly = "soil_temperature_486cm"
	ForecastParamsHourlySoilTemperature54cm              ForecastParamsHourly = "soil_temperature_54cm"
	ForecastParamsHourlySoilTemperature6cm               ForecastParamsHourly = "soil_temperature_6cm"
	ForecastParamsHourlySoilTemperature7To28cm           ForecastParamsHourly = "soil_temperature_7_to_28cm"
	ForecastParamsHourlySunshineDuration                 ForecastParamsHourly = "sunshine_duration"
	ForecastParamsHourlySurfacePressure                  ForecastParamsHourly = "surface_pressure"
	ForecastParamsHourlySurfaceTemperature               ForecastParamsHourly = "surface_temperature"
	ForecastParamsHourlyTemperature1000hPa               ForecastParamsHourly = "temperature_1000hPa"
	ForecastParamsHourlyTemperature100hPa                ForecastParamsHourly = "temperature_100hPa"
	ForecastParamsHourlyTemperature100m                  ForecastParamsHourly = "temperature_100m"
	ForecastParamsHourlyTemperature10hPa                 ForecastParamsHourly = "temperature_10hPa"
	ForecastParamsHourlyTemperature120m                  ForecastParamsHourly = "temperature_120m"
	ForecastParamsHourlyTemperature125hPa                ForecastParamsHourly = "temperature_125hPa"
	ForecastParamsHourlyTemperature150hPa                ForecastParamsHourly = "temperature_150hPa"
	ForecastParamsHourlyTemperature150m                  ForecastParamsHourly = "temperature_150m"
	ForecastParamsHourlyTemperature175hPa                ForecastParamsHourly = "temperature_175hPa"
	ForecastParamsHourlyTemperature180m                  ForecastParamsHourly = "temperature_180m"
	ForecastParamsHourlyTemperature200hPa                ForecastParamsHourly = "temperature_200hPa"
	ForecastParamsHourlyTemperature200m                  ForecastParamsHourly = "temperature_200m"
	ForecastParamsHourlyTemperature20m                   ForecastParamsHourly = "temperature_20m"
	ForecastParamsHourlyTemperature225hPa                ForecastParamsHourly = "temperature_225hPa"
	ForecastParamsHourlyTemperature250hPa                ForecastParamsHourly = "temperature_250hPa"
	ForecastParamsHourlyTemperature275hPa                ForecastParamsHourly = "temperature_275hPa"
	ForecastParamsHourlyTemperature2m                    ForecastParamsHourly = "temperature_2m"
	ForecastParamsHourlyTemperature2mMax                 ForecastParamsHourly = "temperature_2m_max"
	ForecastParamsHourlyTemperature2mMin                 ForecastParamsHourly = "temperature_2m_min"
	ForecastParamsHourlyTemperature300hPa                ForecastParamsHourly = "temperature_300hPa"
	ForecastParamsHourlyTemperature350hPa                ForecastParamsHourly = "temperature_350hPa"
	ForecastParamsHourlyTemperature400hPa                ForecastParamsHourly = "temperature_400hPa"
	ForecastParamsHourlyTemperature40m                   ForecastParamsHourly = "temperature_40m"
	ForecastParamsHourlyTemperature450hPa                ForecastParamsHourly = "temperature_450hPa"
	ForecastParamsHourlyTemperature500hPa                ForecastParamsHourly = "temperature_500hPa"
	ForecastParamsHourlyTemperature50hPa                 ForecastParamsHourly = "temperature_50hPa"
	ForecastParamsHourlyTemperature50m                   ForecastParamsHourly = "temperature_50m"
	ForecastParamsHourlyTemperature550hPa                ForecastParamsHourly = "temperature_550hPa"
	ForecastParamsHourlyTemperature600hPa                ForecastParamsHourly = "temperature_600hPa"
	ForecastParamsHourlyTemperature650hPa                ForecastParamsHourly = "temperature_650hPa"
	ForecastParamsHourlyTemperature700hPa                ForecastParamsHourly = "temperature_700hPa"
	ForecastParamsHourlyTemperature750hPa                ForecastParamsHourly = "temperature_750hPa"
	ForecastParamsHourlyTemperature800hPa                ForecastParamsHourly = "temperature_800hPa"
	ForecastParamsHourlyTemperature80m                   ForecastParamsHourly = "temperature_80m"
	ForecastParamsHourlyTemperature850hPa                ForecastParamsHourly = "temperature_850hPa"
	ForecastParamsHourlyTemperature900hPa                ForecastParamsHourly = "temperature_900hPa"
	ForecastParamsHourlyTemperature925hPa                ForecastParamsHourly = "temperature_925hPa"
	ForecastParamsHourlyTemperature950hPa                ForecastParamsHourly = "temperature_950hPa"
	ForecastParamsHourlyTerrestrialRadiation             ForecastParamsHourly = "terrestrial_radiation"
	ForecastParamsHourlyTerrestrialRadiationInstant      ForecastParamsHourly = "terrestrial_radiation_instant"
	ForecastParamsHourlyThunderstormProbability          ForecastParamsHourly = "thunderstorm_probability"
	ForecastParamsHourlyTotalColumnIntegratedWaterVapour ForecastParamsHourly = "total_column_integrated_water_vapour"
	ForecastParamsHourlyUpdraft                          ForecastParamsHourly = "updraft"
	ForecastParamsHourlyUvIndex                          ForecastParamsHourly = "uv_index"
	ForecastParamsHourlyUvIndexClearSky                  ForecastParamsHourly = "uv_index_clear_sky"
	ForecastParamsHourlyVapourPressureDeficit            ForecastParamsHourly = "vapour_pressure_deficit"
	ForecastParamsHourlyVerticalVelocity1000hPa          ForecastParamsHourly = "vertical_velocity_1000hPa"
	ForecastParamsHourlyVerticalVelocity100hPa           ForecastParamsHourly = "vertical_velocity_100hPa"
	ForecastParamsHourlyVerticalVelocity10hPa            ForecastParamsHourly = "vertical_velocity_10hPa"
	ForecastParamsHourlyVerticalVelocity125hPa           ForecastParamsHourly = "vertical_velocity_125hPa"
	ForecastParamsHourlyVerticalVelocity150hPa           ForecastParamsHourly = "vertical_velocity_150hPa"
	ForecastParamsHourlyVerticalVelocity175hPa           ForecastParamsHourly = "vertical_velocity_175hPa"
	ForecastParamsHourlyVerticalVelocity200hPa           ForecastParamsHourly = "vertical_velocity_200hPa"
	ForecastParamsHourlyVerticalVelocity225hPa           ForecastParamsHourly = "vertical_velocity_225hPa"
	ForecastParamsHourlyVerticalVelocity250hPa           ForecastParamsHourly = "vertical_velocity_250hPa"
	ForecastParamsHourlyVerticalVelocity275hPa           ForecastParamsHourly = "vertical_velocity_275hPa"
	ForecastParamsHourlyVerticalVelocity300hPa           ForecastParamsHourly = "vertical_velocity_300hPa"
	ForecastParamsHourlyVerticalVelocity350hPa           ForecastParamsHourly = "vertical_velocity_350hPa"
	ForecastParamsHourlyVerticalVelocity400hPa           ForecastParamsHourly = "vertical_velocity_400hPa"
	ForecastParamsHourlyVerticalVelocity450hPa           ForecastParamsHourly = "vertical_velocity_450hPa"
	ForecastParamsHourlyVerticalVelocity500hPa           ForecastParamsHourly = "vertical_velocity_500hPa"
	ForecastParamsHourlyVerticalVelocity50hPa            ForecastParamsHourly = "vertical_velocity_50hPa"
	ForecastParamsHourlyVerticalVelocity550hPa           ForecastParamsHourly = "vertical_velocity_550hPa"
	ForecastParamsHourlyVerticalVelocity600hPa           ForecastParamsHourly = "vertical_velocity_600hPa"
	ForecastParamsHourlyVerticalVelocity650hPa           ForecastParamsHourly = "vertical_velocity_650hPa"
	ForecastParamsHourlyVerticalVelocity700hPa           ForecastParamsHourly = "vertical_velocity_700hPa"
	ForecastParamsHourlyVerticalVelocity750hPa           ForecastParamsHourly = "vertical_velocity_750hPa"
	ForecastParamsHourlyVerticalVelocity800hPa           ForecastParamsHourly = "vertical_velocity_800hPa"
	ForecastParamsHourlyVerticalVelocity850hPa           ForecastParamsHourly = "vertical_velocity_850hPa"
	ForecastParamsHourlyVerticalVelocity900hPa           ForecastParamsHourly = "vertical_velocity_900hPa"
	ForecastParamsHourlyVerticalVelocity925hPa           ForecastParamsHourly = "vertical_velocity_925hPa"
	ForecastParamsHourlyVerticalVelocity950hPa           ForecastParamsHourly = "vertical_velocity_950hPa"
	ForecastParamsHourlyVisibility                       ForecastParamsHourly = "visibility"
	ForecastParamsHourlyWeatherCode                      ForecastParamsHourly = "weather_code"
	ForecastParamsHourlyWetBulbTemperature2m             ForecastParamsHourly = "wet_bulb_temperature_2m"
	ForecastParamsHourlyWindDirection1000hPa             ForecastParamsHourly = "wind_direction_1000hPa"
	ForecastParamsHourlyWindDirection100hPa              ForecastParamsHourly = "wind_direction_100hPa"
	ForecastParamsHourlyWindDirection100m                ForecastParamsHourly = "wind_direction_100m"
	ForecastParamsHourlyWindDirection10hPa               ForecastParamsHourly = "wind_direction_10hPa"
	ForecastParamsHourlyWindDirection10m                 ForecastParamsHourly = "wind_direction_10m"
	ForecastParamsHourlyWindDirection120m                ForecastParamsHourly = "wind_direction_120m"
	ForecastParamsHourlyWindDirection125hPa              ForecastParamsHourly = "wind_direction_125hPa"
	ForecastParamsHourlyWindDirection140m                ForecastParamsHourly = "wind_direction_140m"
	ForecastParamsHourlyWindDirection150hPa              ForecastParamsHourly = "wind_direction_150hPa"
	ForecastParamsHourlyWindDirection150m                ForecastParamsHourly = "wind_direction_150m"
	ForecastParamsHourlyWindDirection160m                ForecastParamsHourly = "wind_direction_160m"
	ForecastParamsHourlyWindDirection175hPa              ForecastParamsHourly = "wind_direction_175hPa"
	ForecastParamsHourlyWindDirection180m                ForecastParamsHourly = "wind_direction_180m"
	ForecastParamsHourlyWindDirection200hPa              ForecastParamsHourly = "wind_direction_200hPa"
	ForecastParamsHourlyWindDirection200m                ForecastParamsHourly = "wind_direction_200m"
	ForecastParamsHourlyWindDirection20m                 ForecastParamsHourly = "wind_direction_20m"
	ForecastParamsHourlyWindDirection225hPa              ForecastParamsHourly = "wind_direction_225hPa"
	ForecastParamsHourlyWindDirection250hPa              ForecastParamsHourly = "wind_direction_250hPa"
	ForecastParamsHourlyWindDirection275hPa              ForecastParamsHourly = "wind_direction_275hPa"
	ForecastParamsHourlyWindDirection300hPa              ForecastParamsHourly = "wind_direction_300hPa"
	ForecastParamsHourlyWindDirection30m                 ForecastParamsHourly = "wind_direction_30m"
	ForecastParamsHourlyWindDirection350hPa              ForecastParamsHourly = "wind_direction_350hPa"
	ForecastParamsHourlyWindDirection400hPa              ForecastParamsHourly = "wind_direction_400hPa"
	ForecastParamsHourlyWindDirection40m                 ForecastParamsHourly = "wind_direction_40m"
	ForecastParamsHourlyWindDirection450hPa              ForecastParamsHourly = "wind_direction_450hPa"
	ForecastParamsHourlyWindDirection500hPa              ForecastParamsHourly = "wind_direction_500hPa"
	ForecastParamsHourlyWindDirection50hPa               ForecastParamsHourly = "wind_direction_50hPa"
	ForecastParamsHourlyWindDirection50m                 ForecastParamsHourly = "wind_direction_50m"
	ForecastParamsHourlyWindDirection550hPa              ForecastParamsHourly = "wind_direction_550hPa"
	ForecastParamsHourlyWindDirection600hPa              ForecastParamsHourly = "wind_direction_600hPa"
	ForecastParamsHourlyWindDirection650hPa              ForecastParamsHourly = "wind_direction_650hPa"
	ForecastParamsHourlyWindDirection700hPa              ForecastParamsHourly = "wind_direction_700hPa"
	ForecastParamsHourlyWindDirection70m                 ForecastParamsHourly = "wind_direction_70m"
	ForecastParamsHourlyWindDirection750hPa              ForecastParamsHourly = "wind_direction_750hPa"
	ForecastParamsHourlyWindDirection800hPa              ForecastParamsHourly = "wind_direction_800hPa"
	ForecastParamsHourlyWindDirection80m                 ForecastParamsHourly = "wind_direction_80m"
	ForecastParamsHourlyWindDirection850hPa              ForecastParamsHourly = "wind_direction_850hPa"
	ForecastParamsHourlyWindDirection900hPa              ForecastParamsHourly = "wind_direction_900hPa"
	ForecastParamsHourlyWindDirection925hPa              ForecastParamsHourly = "wind_direction_925hPa"
	ForecastParamsHourlyWindDirection950hPa              ForecastParamsHourly = "wind_direction_950hPa"
	ForecastParamsHourlyWindGusts10m                     ForecastParamsHourly = "wind_gusts_10m"
	ForecastParamsHourlyWindSpeed1000hPa                 ForecastParamsHourly = "wind_speed_1000hPa"
	ForecastParamsHourlyWindSpeed100hPa                  ForecastParamsHourly = "wind_speed_100hPa"
	ForecastParamsHourlyWindSpeed100m                    ForecastParamsHourly = "wind_speed_100m"
	ForecastParamsHourlyWindSpeed10hPa                   ForecastParamsHourly = "wind_speed_10hPa"
	ForecastParamsHourlyWindSpeed10m                     ForecastParamsHourly = "wind_speed_10m"
	ForecastParamsHourlyWindSpeed120m                    ForecastParamsHourly = "wind_speed_120m"
	ForecastParamsHourlyWindSpeed125hPa                  ForecastParamsHourly = "wind_speed_125hPa"
	ForecastParamsHourlyWindSpeed140m                    ForecastParamsHourly = "wind_speed_140m"
	ForecastParamsHourlyWindSpeed150hPa                  ForecastParamsHourly = "wind_speed_150hPa"
	ForecastParamsHourlyWindSpeed150m                    ForecastParamsHourly = "wind_speed_150m"
	ForecastParamsHourlyWindSpeed160m                    ForecastParamsHourly = "wind_speed_160m"
	ForecastParamsHourlyWindSpeed175hPa                  ForecastParamsHourly = "wind_speed_175hPa"
	ForecastParamsHourlyWindSpeed180m                    ForecastParamsHourly = "wind_speed_180m"
	ForecastParamsHourlyWindSpeed200hPa                  ForecastParamsHourly = "wind_speed_200hPa"
	ForecastParamsHourlyWindSpeed200m                    ForecastParamsHourly = "wind_speed_200m"
	ForecastParamsHourlyWindSpeed20m                     ForecastParamsHourly = "wind_speed_20m"
	ForecastParamsHourlyWindSpeed225hPa                  ForecastParamsHourly = "wind_speed_225hPa"
	ForecastParamsHourlyWindSpeed250hPa                  ForecastParamsHourly = "wind_speed_250hPa"
	ForecastParamsHourlyWindSpeed275hPa                  ForecastParamsHourly = "wind_speed_275hPa"
	ForecastParamsHourlyWindSpeed300hPa                  ForecastParamsHourly = "wind_speed_300hPa"
	ForecastParamsHourlyWindSpeed30m                     ForecastParamsHourly = "wind_speed_30m"
	ForecastParamsHourlyWindSpeed350hPa                  ForecastParamsHourly = "wind_speed_350hPa"
	ForecastParamsHourlyWindSpeed400hPa                  ForecastParamsHourly = "wind_speed_400hPa"
	ForecastParamsHourlyWindSpeed40m                     ForecastParamsHourly = "wind_speed_40m"
	ForecastParamsHourlyWindSpeed450hPa                  ForecastParamsHourly = "wind_speed_450hPa"
	ForecastParamsHourlyWindSpeed500hPa                  ForecastParamsHourly = "wind_speed_500hPa"
	ForecastParamsHourlyWindSpeed50hPa                   ForecastParamsHourly = "wind_speed_50hPa"
	ForecastParamsHourlyWindSpeed50m                     ForecastParamsHourly = "wind_speed_50m"
	ForecastParamsHourlyWindSpeed550hPa                  ForecastParamsHourly = "wind_speed_550hPa"
	ForecastParamsHourlyWindSpeed600hPa                  ForecastParamsHourly = "wind_speed_600hPa"
	ForecastParamsHourlyWindSpeed650hPa                  ForecastParamsHourly = "wind_speed_650hPa"
	ForecastParamsHourlyWindSpeed700hPa                  ForecastParamsHourly = "wind_speed_700hPa"
	ForecastParamsHourlyWindSpeed70m                     ForecastParamsHourly = "wind_speed_70m"
	ForecastParamsHourlyWindSpeed750hPa                  ForecastParamsHourly = "wind_speed_750hPa"
	ForecastParamsHourlyWindSpeed800hPa                  ForecastParamsHourly = "wind_speed_800hPa"
	ForecastParamsHourlyWindSpeed80m                     ForecastParamsHourly = "wind_speed_80m"
	ForecastParamsHourlyWindSpeed850hPa                  ForecastParamsHourly = "wind_speed_850hPa"
	ForecastParamsHourlyWindSpeed900hPa                  ForecastParamsHourly = "wind_speed_900hPa"
	ForecastParamsHourlyWindSpeed925hPa                  ForecastParamsHourly = "wind_speed_925hPa"
	ForecastParamsHourlyWindSpeed950hPa                  ForecastParamsHourly = "wind_speed_950hPa"
)

// Defines values for ForecastParamsDaily.
const (
	ForecastParamsDailyApparentTemperatureMax      ForecastParamsDaily = "apparent_temperature_max"
	ForecastParamsDailyApparentTemperatureMin      ForecastParamsDaily = "apparent_temperature_min"
	ForecastParamsDailyDaylightDuration            ForecastParamsDaily = "daylight_duration"
	ForecastParamsDailyEt0FaoEvapotranspiration    ForecastParamsDaily = "et0_fao_evapotranspiration"
	ForecastParamsDailyPrecipitationHours          ForecastParamsDaily = "precipitation_hours"
	ForecastParamsDailyPrecipitationProbabilityMax ForecastParamsDaily = "precipitation_probability_max"
	ForecastParamsDailyPrecipitationSum            ForecastParamsDaily = "precipitation_sum"
	ForecastParamsDailyRainSum                     ForecastParamsDaily = "rain_sum"
	ForecastParamsDailyShortwaveRadiationSum       ForecastParamsDaily = "shortwave_radiation_sum"
	ForecastParamsDailyShowersSum                  ForecastParamsDaily = "showers_sum"
	ForecastParamsDailySnowfallSum                 ForecastParamsDaily = "snowfall_sum"
	ForecastParamsDailySunrise                     ForecastParamsDaily = "sunrise"
	ForecastParamsDailySunset                      ForecastParamsDaily = "sunset"
	ForecastParamsDailySunshineDuration            ForecastParamsDaily = "sunshine_duration"
	ForecastParamsDailyTemperature2mMax            ForecastParamsDaily = "temperature_2m_max"
	ForecastParamsDailyTemperature2mMin            ForecastParamsDaily = "temperature_2m_min"
	ForecastParamsDailyUvIndexClearSkyMax          ForecastParamsDaily = "uv_index_clear_sky_max"
	ForecastParamsDailyUvIndexMax                  ForecastParamsDaily = "uv_index_max"
	ForecastParamsDailyWeatherCode                 ForecastParamsDaily = "weather_code"
	ForecastParamsDailyWindDirection10mDominant    ForecastParamsDaily = "wind_direction_10m_dominant"
	ForecastParamsDailyWindGusts10mMax             ForecastParamsDaily = "wind_gusts_10m_max"
	ForecastParamsDailyWindSpeed10mMax             ForecastParamsDaily = "wind_speed_10m_max"
)

// Defines values for ForecastParamsCurrent.
const (
	ForecastParamsCurrentApparentTemperature ForecastParamsCurrent = "apparent_temperature"
	ForecastParamsCurrentCloudCover          ForecastParamsCurrent = "cloud_cover"
	ForecastParamsCurrentIsDay               ForecastParamsCurrent = "is_day"
	ForecastParamsCurrentPrecipitation       ForecastParamsCurrent = "precipitation"
	ForecastParamsCurrentPressureMsl         ForecastParamsCurrent = "pressure_msl"
	ForecastParamsCurrentRain                ForecastParamsCurrent = "rain"
	ForecastParamsCurrentRelativeHumidity2m  ForecastParamsCurrent = "relative_humidity_2m"
	ForecastParamsCurrentShowers             ForecastParamsCurrent = "showers"
	ForecastParamsCurrentSnowfall            ForecastParamsCurrent = "snowfall"
	ForecastParamsCurrentSurfacePressure     ForecastParamsCurrent = "surface_pressure"
	ForecastParamsCurrentTemperature2m       ForecastParamsCurrent = "temperature_2m"
	ForecastParamsCurrentWeatherCode         ForecastParamsCurrent = "weather_code"
	ForecastParamsCurrentWindDirection10m    ForecastParamsCurrent = "wind_direction_10m"
	ForecastParamsCurrentWindGusts10m        ForecastParamsCurrent = "wind_gusts_10m"
	ForecastParamsCurrentWindSpeed10m        ForecastParamsCurrent = "wind_speed_10m"
)

// Defines values for ForecastParamsMinutely15.
const (
	ForecastParamsMinutely15ApparentTemperature ForecastParamsMinutely15 = "apparent_temperature"
	ForecastParamsMinutely15Cape                ForecastParamsMinutely15 = "cape"
	ForecastParamsMinutely15DewPoint2m          ForecastParamsMinutely15 = "dew_point_2m"
	ForecastParamsMinutely15FreezingLevelHeight ForecastParamsMinutely15 = "freezing_level_height"
	ForecastParamsMinutely15IsDay               ForecastParamsMinutely15 = "is_day"
	ForecastParamsMinutely15LightningPotential  ForecastParamsMinutely15 = "lightning_potential"
	ForecastParamsMinutely15Precipitation       ForecastParamsMinutely15 = "precipitation"
	ForecastParamsMinutely15Rain                ForecastParamsMinutely15 = "rain"
	ForecastParamsMinutely15RelativeHumidity2m  ForecastParamsMinutely15 = "relative_humidity_2m"
	ForecastParamsMinutely15Snowfall            ForecastParamsMinutely15 = "snowfall"
	ForecastParamsMinutely15SnowfallHeight      ForecastParamsMinutely15 = "snowfall_height"
	ForecastParamsMinutely15SunshineDuration    ForecastParamsMinutely15 = "sunshine_duration"
	ForecastParamsMinutely15Temperature2m       ForecastParamsMinutely15 = "temperature_2m"
	ForecastParamsMinutely15Visibility          ForecastParamsMinutely15 = "visibility"
	ForecastParamsMinutely15WeatherCode         ForecastParamsMinutely15 = "weather_code"
	ForecastParamsMinutely15WindDirection100m   ForecastParamsMinutely15 = "wind_direction_100m"
	ForecastParamsMinutely15WindDirection10m    ForecastParamsMinutely15 = "wind_direction_10m"
	ForecastParamsMinutely15WindDirection20m    ForecastParamsMinutely15 = "wind_direction_20m"
	ForecastParamsMinutely15WindDirection50m    ForecastParamsMinutely15 = "wind_direction_50m"
	ForecastParamsMinutely15WindDirection80m    ForecastParamsMinutely15 = "wind_direction_80m"
	ForecastParamsMinutely15WindGusts10m        ForecastParamsMinutely15 = "wind_gusts_10m"
	ForecastParamsMinutely15WindSpeed100m       ForecastParamsMinutely15 = "wind_speed_100m"
	ForecastParamsMinutely15WindSpeed10m        ForecastParamsMinutely15 = "wind_speed_10m"
	ForecastParamsMinutely15WindSpeed20m        ForecastParamsMinutely15 = "wind_speed_20m"
	ForecastParamsMinutely15WindSpeed50m        ForecastParamsMinutely15 = "wind_speed_50m"
	ForecastParamsMinutely15WindSpeed80m        ForecastParamsMinutely15 = "wind_speed_80m"
)

// Defines values for ForecastParamsTemperatureUnit.
const (
	Celsius    ForecastParamsTemperatureUnit = "celsius"
	Fahrenheit ForecastParamsTemperatureUnit = "fahrenheit"
)

// Defines values for ForecastParamsWindSpeedUnit.
const (
	Kmh ForecastParamsWindSpeedUnit = "kmh"
	Kn  ForecastParamsWindSpeedUnit = "kn"
	Mph ForecastParamsWindSpeedUnit = "mph"
	Ms  ForecastParamsWindSpeedUnit = "ms"
)

const (
	Inch ForecastParamsPrecipitationUnit = "inch"
	Mm   ForecastParamsPrecipitationUnit = "mm"
)

const (
	Iso8601  ForecastParamsTimeformat = "iso8601"
	Unixtime ForecastParamsTimeformat = "unixtime"
)

// Defines values for ForecastParamsCellSelection.
const (
	Land    ForecastParamsCellSelection = "land"
	Nearest ForecastParamsCellSelection = "nearest"
	Sea     ForecastParamsCellSelection = "sea"
)

// Defines values for ForecastParamsModels.
const (
	BestMatch                    ForecastParamsModels = "best_match"
	BomAccessGlobal              ForecastParamsModels = "bom_access_global"
	CmaGrapesGlobal              ForecastParamsModels = "cma_grapes_global"
	CmcGemGdps                   ForecastParamsModels = "cmc_gem_gdps"
	CmcGemHrdps                  ForecastParamsModels = "cmc_gem_hrdps"
	CmcGemHrdpsWest              ForecastParamsModels = "cmc_gem_hrdps_west"
	CmcGemRdps                   ForecastParamsModels = "cmc_gem_rdps"
	CmcGemSeamless               ForecastParamsModels = "cmc_gem_seamless"
	DmiHarmonieAromeEurope       ForecastParamsModels = "dmi_harmonie_arome_europe"
	DmiSeamless                  ForecastParamsModels = "dmi_seamless"
	EcmwfAifs025Single           ForecastParamsModels = "ecmwf_aifs025_single"
	EcmwfIfs                     ForecastParamsModels = "ecmwf_ifs"
	EcmwfIfs025                  ForecastParamsModels = "ecmwf_ifs025"
	GeosphereAromeAustria        ForecastParamsModels = "geosphere_arome_austria"
	GeosphereSeamless            ForecastParamsModels = "geosphere_seamless"
	IconD2                       ForecastParamsModels = "icon_d2"
	IconEu                       ForecastParamsModels = "icon_eu"
	IconGlobal                   ForecastParamsModels = "icon_global"
	IconSeamless                 ForecastParamsModels = "icon_seamless"
	ItaliaMeteoArpaeIcon2i       ForecastParamsModels = "italia_meteo_arpae_icon_2i"
	JmaGsm                       ForecastParamsModels = "jma_gsm"
	JmaMsm                       ForecastParamsModels = "jma_msm"
	JmaSeamless                  ForecastParamsModels = "jma_seamless"
	KmaGdps                      ForecastParamsModels = "kma_gdps"
	KmaLdps                      ForecastParamsModels = "kma_ldps"
	KmaSeamless                  ForecastParamsModels = "kma_seamless"
	KnmiHarmonieAromeEurope      ForecastParamsModels = "knmi_harmonie_arome_europe"
	KnmiHarmonieAromeNetherlands ForecastParamsModels = "knmi_harmonie_arome_netherlands"
	KnmiSeamless                 ForecastParamsModels = "knmi_seamless"
	MeteofranceAromeFrance       ForecastParamsModels = "meteofrance_arome_france"
	MeteofranceAromeFranceHd     ForecastParamsModels = "meteofrance_arome_france_hd"
	MeteofranceArpegeEurope      ForecastParamsModels = "meteofrance_arpege_europe"
	MeteofranceArpegeWorld       ForecastParamsModels = "meteofrance_arpege_world"
	MeteofranceSeamless          ForecastParamsModels = "meteofrance_seamless"
	MeteoswissIconCh1            ForecastParamsModels = "meteoswiss_icon_ch1"
	MeteoswissIconCh2            ForecastParamsModels = "meteoswiss_icon_ch2"
	MeteoswissIconSeamless       ForecastParamsModels = "meteoswiss_icon_seamless"
	MetnoNordic                  ForecastParamsModels = "metno_nordic"
	MetnoSeamless                ForecastParamsModels = "metno_seamless"
	NcepAigfs025                 ForecastParamsModels = "ncep_aigfs025"
	NcepGfsGlobal                ForecastParamsModels = "ncep_gfs_global"
	NcepGfsGraphcast025          ForecastParamsModels = "ncep_gfs_graphcast025"
	NcepGfsSeamless              ForecastParamsModels = "ncep_gfs_seamless"
	NcepHgefs025EnsembleMean     ForecastParamsModels = "ncep_hgefs025_ensemble_mean"
	NcepHrrrConus                ForecastParamsModels = "ncep_hrrr_conus"
	NcepNamConus                 ForecastParamsModels = "ncep_nam_conus"
	NcepNbmConus                 ForecastParamsModels = "ncep_nbm_conus"
	UkmoGlobalDeterministic10km  ForecastParamsModels = "ukmo_global_deterministic_10km"
	UkmoSeamless                 ForecastParamsModels = "ukmo_seamless"
	UkmoUkDeterministic2km       ForecastParamsModels = "ukmo_uk_deterministic_2km"
)

type ForecastRequest struct {
	Latitude          float64                         `form:"latitude" json:"latitude"`
	Longitude         float64                         `form:"longitude" json:"longitude"`
	Hourly            []ForecastParamsHourly          `form:"hourly" json:"hourly"`
	Daily             []ForecastParamsDaily           `form:"daily" json:"daily"`
	Current           []ForecastParamsCurrent         `form:"current" json:"current"`
	Minutely15        []ForecastParamsMinutely15      `form:"minutely_15" json:"minutely_15"`
	Elevation         float32                         `form:"elevation" json:"elevation"`
	TemperatureUnit   ForecastParamsTemperatureUnit   `form:"temperature_unit" json:"temperature_unit"`
	WindSpeedUnit     ForecastParamsWindSpeedUnit     `form:"wind_speed_unit" json:"wind_speed_unit"`
	PrecipitationUnit ForecastParamsPrecipitationUnit `form:"precipitation_unit" json:"precipitation_unit"`
	Timeformat        ForecastParamsTimeformat        `form:"timeformat" json:"timeformat"`
	Timezone          string                          `form:"timezone" json:"timezone"`
	PastDays          int                             `form:"past_days" json:"past_days"`
	ForecastDays      int                             `form:"forecast_days" json:"forecast_days"`
	PastHours         int                             `form:"past_hours" json:"past_hours"`
	ForecastHours     int                             `form:"forecast_hours" json:"forecast_hours"`
	StartDate         string                          `form:"start_date" json:"start_date"`
	EndDate           string                          `form:"end_date" json:"end_date"`
	Tilt              float32                         `form:"tilt" json:"tilt"`
	Azimuth           float32                         `form:"azimuth" json:"azimuth"`
	CellSelection     ForecastParamsCellSelection     `form:"cell_selection" json:"cell_selection"`
	Apikey            string                          `form:"apikey" json:"apikey"`
	Models            []ForecastParamsModels          `form:"models" json:"models"`
}

// ForecastParamsHourly defines parameters for Forecast.
type ForecastParamsHourly string

// ForecastParamsDaily defines parameters for Forecast.
type ForecastParamsDaily string

// ForecastParamsCurrent defines parameters for Forecast.
type ForecastParamsCurrent string

// ForecastParamsMinutely15 defines parameters for Forecast.
type ForecastParamsMinutely15 string

// ForecastParamsTemperatureUnit defines parameters for Forecast.
type ForecastParamsTemperatureUnit string

// ForecastParamsWindSpeedUnit defines parameters for Forecast.
type ForecastParamsWindSpeedUnit string

// ForecastParamsPrecipitationUnit defines parameters for Forecast.
type ForecastParamsPrecipitationUnit string

// ForecastParamsTimeformat defines parameters for Forecast.
type ForecastParamsTimeformat string

// ForecastParamsCellSelection defines parameters for Forecast.
type ForecastParamsCellSelection string

// ForecastParamsModels defines parameters for Forecast.
type ForecastParamsModels string

type Current struct {
	ApparentTemperature float32 `json:"apparent_temperature"`
	CloudCover          float32 `json:"cloud_cover"`
	Interval            int     `json:"interval"`
	IsDay               int     `json:"is_day"`
	Precipitation       float32 `json:"precipitation"`
	PressureMsl         float32 `json:"pressure_msl"`
	Rain                float32 `json:"rain"`
	RelativeHumidity2m  float32 `json:"relative_humidity_2m"`
	Showers             float32 `json:"showers"`
	Snowfall            float32 `json:"snowfall"`
	SurfacePressure     float32 `json:"surface_pressure"`
	Temperature2m       float32 `json:"temperature_2m"`
	Time                string  `json:"time"`
	WeatherCode         int     `json:"weather_code"`
	WindDirection10m    float32 `json:"wind_direction_10m"`
	WindGusts10m        float32 `json:"wind_gusts_10m"`
	WindSpeed10m        float32 `json:"wind_speed_10m"`
}

type CurrentUnits struct {
	ApparentTemperature string `json:"apparent_temperature"`
	CloudCover          string `json:"cloud_cover"`
	IsDay               string `json:"is_day"`
	Precipitation       string `json:"precipitation"`
	PressureMsl         string `json:"pressure_msl"`
	Rain                string `json:"rain"`
	RelativeHumidity2m  string `json:"relative_humidity_2m"`
	Showers             string `json:"showers"`
	Snowfall            string `json:"snowfall"`
	SurfacePressure     string `json:"surface_pressure"`
	Temperature2m       string `json:"temperature_2m"`
	Time                string `json:"time"`
	WeatherCode         string `json:"weather_code"`
	WindDirection10m    string `json:"wind_direction_10m"`
	WindGusts10m        string `json:"wind_gusts_10m"`
	WindSpeed10m        string `json:"wind_speed_10m"`
}

type Daily struct {
	ApparentTemperatureMax      []float32 `json:"apparent_temperature_max"`
	ApparentTemperatureMin      []float32 `json:"apparent_temperature_min"`
	DaylightDuration            []float32 `json:"daylight_duration"`
	Et0FaoEvapotranspiration    []float32 `json:"et0_fao_evapotranspiration"`
	PrecipitationHours          []float32 `json:"precipitation_hours"`
	PrecipitationProbabilityMax []float32 `json:"precipitation_probability_max"`
	PrecipitationSum            []float32 `json:"precipitation_sum"`
	RainSum                     []float32 `json:"rain_sum"`
	ShortwaveRadiationSum       []float32 `json:"shortwave_radiation_sum"`
	ShowersSum                  []float32 `json:"showers_sum"`
	SnowfallSum                 []float32 `json:"snowfall_sum"`
	Sunrise                     []string  `json:"sunrise"`
	Sunset                      []string  `json:"sunset"`
	SunshineDuration            []float32 `json:"sunshine_duration"`
	Temperature2mMax            []float32 `json:"temperature_2m_max"`
	Temperature2mMin            []float32 `json:"temperature_2m_min"`
	Time                        []string  `json:"time"`
	UvIndexClearSkyMax          []float32 `json:"uv_index_clear_sky_max"`
	UvIndexMax                  []float32 `json:"uv_index_max"`
	WeatherCode                 []int     `json:"weather_code"`
	WindDirection10mDominant    []float32 `json:"wind_direction_10m_dominant"`
	WindGusts10mMax             []float32 `json:"wind_gusts_10m_max"`
	WindSpeed10mMax             []float32 `json:"wind_speed_10m_max"`
}

type DailyUnits struct {
	ApparentTemperatureMax      string `json:"apparent_temperature_max"`
	ApparentTemperatureMin      string `json:"apparent_temperature_min"`
	DaylightDuration            string `json:"daylight_duration"`
	Et0FaoEvapotranspiration    string `json:"et0_fao_evapotranspiration"`
	PrecipitationHours          string `json:"precipitation_hours"`
	PrecipitationProbabilityMax string `json:"precipitation_probability_max"`
	PrecipitationSum            string `json:"precipitation_sum"`
	RainSum                     string `json:"rain_sum"`
	ShortwaveRadiationSum       string `json:"shortwave_radiation_sum"`
	ShowersSum                  string `json:"showers_sum"`
	SnowfallSum                 string `json:"snowfall_sum"`
	Sunrise                     string `json:"sunrise"`
	Sunset                      string `json:"sunset"`
	SunshineDuration            string `json:"sunshine_duration"`
	Temperature2mMax            string `json:"temperature_2m_max"`
	Temperature2mMin            string `json:"temperature_2m_min"`
	Time                        string `json:"time"`
	UvIndexClearSkyMax          string `json:"uv_index_clear_sky_max"`
	UvIndexMax                  string `json:"uv_index_max"`
	WeatherCode                 string `json:"weather_code"`
	WindDirection10mDominant    string `json:"wind_direction_10m_dominant"`
	WindGusts10mMax             string `json:"wind_gusts_10m_max"`
	WindSpeed10mMax             string `json:"wind_speed_10m_max"`
}

type Hourly struct {
	Albedo                           []float32 `json:"albedo"`
	ApparentTemperature              []float32 `json:"apparent_temperature"`
	BoundaryLayerHeight              []float32 `json:"boundary_layer_height"`
	Cape                             []float32 `json:"cape"`
	CloudCover                       []float32 `json:"cloud_cover"`
	CloudCover1000hPa                []float32 `json:"cloud_cover_1000hPa"`
	CloudCover100hPa                 []float32 `json:"cloud_cover_100hPa"`
	CloudCover10hPa                  []float32 `json:"cloud_cover_10hPa"`
	CloudCover125hPa                 []float32 `json:"cloud_cover_125hPa"`
	CloudCover150hPa                 []float32 `json:"cloud_cover_150hPa"`
	CloudCover175hPa                 []float32 `json:"cloud_cover_175hPa"`
	CloudCover200hPa                 []float32 `json:"cloud_cover_200hPa"`
	CloudCover225hPa                 []float32 `json:"cloud_cover_225hPa"`
	CloudCover250hPa                 []float32 `json:"cloud_cover_250hPa"`
	CloudCover275hPa                 []float32 `json:"cloud_cover_275hPa"`
	CloudCover300hPa                 []float32 `json:"cloud_cover_300hPa"`
	CloudCover350hPa                 []float32 `json:"cloud_cover_350hPa"`
	CloudCover400hPa                 []float32 `json:"cloud_cover_400hPa"`
	CloudCover450hPa                 []float32 `json:"cloud_cover_450hPa"`
	CloudCover500hPa                 []float32 `json:"cloud_cover_500hPa"`
	CloudCover50hPa                  []float32 `json:"cloud_cover_50hPa"`
	CloudCover550hPa                 []float32 `json:"cloud_cover_550hPa"`
	CloudCover600hPa                 []float32 `json:"cloud_cover_600hPa"`
	CloudCover650hPa                 []float32 `json:"cloud_cover_650hPa"`
	CloudCover700hPa                 []float32 `json:"cloud_cover_700hPa"`
	CloudCover750hPa                 []float32 `json:"cloud_cover_750hPa"`
	CloudCover800hPa                 []float32 `json:"cloud_cover_800hPa"`
	CloudCover850hPa                 []float32 `json:"cloud_cover_850hPa"`
	CloudCover900hPa                 []float32 `json:"cloud_cover_900hPa"`
	CloudCover925hPa                 []float32 `json:"cloud_cover_925hPa"`
	CloudCover950hPa                 []float32 `json:"cloud_cover_950hPa"`
	CloudCoverHigh                   []float32 `json:"cloud_cover_high"`
	CloudCoverLow                    []float32 `json:"cloud_cover_low"`
	CloudCoverMid                    []float32 `json:"cloud_cover_mid"`
	ConvectiveCloudBase              []float32 `json:"convective_cloud_base"`
	ConvectiveCloudTop               []float32 `json:"convective_cloud_top"`
	ConvectiveInhibition             []float32 `json:"convective_inhibition"`
	DewPoint1000hPa                  []float32 `json:"dew_point_1000hPa"`
	DewPoint100hPa                   []float32 `json:"dew_point_100hPa"`
	DewPoint10hPa                    []float32 `json:"dew_point_10hPa"`
	DewPoint125hPa                   []float32 `json:"dew_point_125hPa"`
	DewPoint150hPa                   []float32 `json:"dew_point_150hPa"`
	DewPoint175hPa                   []float32 `json:"dew_point_175hPa"`
	DewPoint200hPa                   []float32 `json:"dew_point_200hPa"`
	DewPoint225hPa                   []float32 `json:"dew_point_225hPa"`
	DewPoint250hPa                   []float32 `json:"dew_point_250hPa"`
	DewPoint275hPa                   []float32 `json:"dew_point_275hPa"`
	DewPoint2m                       []float32 `json:"dew_point_2m"`
	DewPoint300hPa                   []float32 `json:"dew_point_300hPa"`
	DewPoint350hPa                   []float32 `json:"dew_point_350hPa"`
	DewPoint400hPa                   []float32 `json:"dew_point_400hPa"`
	DewPoint450hPa                   []float32 `json:"dew_point_450hPa"`
	DewPoint500hPa                   []float32 `json:"dew_point_500hPa"`
	DewPoint50hPa                    []float32 `json:"dew_point_50hPa"`
	DewPoint550hPa                   []float32 `json:"dew_point_550hPa"`
	DewPoint600hPa                   []float32 `json:"dew_point_600hPa"`
	DewPoint650hPa                   []float32 `json:"dew_point_650hPa"`
	DewPoint700hPa                   []float32 `json:"dew_point_700hPa"`
	DewPoint750hPa                   []float32 `json:"dew_point_750hPa"`
	DewPoint800hPa                   []float32 `json:"dew_point_800hPa"`
	DewPoint850hPa                   []float32 `json:"dew_point_850hPa"`
	DewPoint900hPa                   []float32 `json:"dew_point_900hPa"`
	DewPoint925hPa                   []float32 `json:"dew_point_925hPa"`
	DewPoint950hPa                   []float32 `json:"dew_point_950hPa"`
	DiffuseRadiation                 []float32 `json:"diffuse_radiation"`
	DiffuseRadiationInstant          []float32 `json:"diffuse_radiation_instant"`
	DirectNormalIrradiance           []float32 `json:"direct_normal_irradiance"`
	DirectNormalIrradianceInstant    []float32 `json:"direct_normal_irradiance_instant"`
	DirectRadiation                  []float32 `json:"direct_radiation"`
	DirectRadiationInstant           []float32 `json:"direct_radiation_instant"`
	Et0FaoEvapotranspiration         []float32 `json:"et0_fao_evapotranspiration"`
	Evapotranspiration               []float32 `json:"evapotranspiration"`
	FreezingLevelHeight              []float32 `json:"freezing_level_height"`
	FreezingRainProbability          []float32 `json:"freezing_rain_probability"`
	GeopotentialHeight1000hPa        []float32 `json:"geopotential_height_1000hPa"`
	GeopotentialHeight100hPa         []float32 `json:"geopotential_height_100hPa"`
	GeopotentialHeight10hPa          []float32 `json:"geopotential_height_10hPa"`
	GeopotentialHeight125hPa         []float32 `json:"geopotential_height_125hPa"`
	GeopotentialHeight150hPa         []float32 `json:"geopotential_height_150hPa"`
	GeopotentialHeight175hPa         []float32 `json:"geopotential_height_175hPa"`
	GeopotentialHeight200hPa         []float32 `json:"geopotential_height_200hPa"`
	GeopotentialHeight225hPa         []float32 `json:"geopotential_height_225hPa"`
	GeopotentialHeight250hPa         []float32 `json:"geopotential_height_250hPa"`
	GeopotentialHeight275hPa         []float32 `json:"geopotential_height_275hPa"`
	GeopotentialHeight300hPa         []float32 `json:"geopotential_height_300hPa"`
	GeopotentialHeight350hPa         []float32 `json:"geopotential_height_350hPa"`
	GeopotentialHeight400hPa         []float32 `json:"geopotential_height_400hPa"`
	GeopotentialHeight450hPa         []float32 `json:"geopotential_height_450hPa"`
	GeopotentialHeight500hPa         []float32 `json:"geopotential_height_500hPa"`
	GeopotentialHeight50hPa          []float32 `json:"geopotential_height_50hPa"`
	GeopotentialHeight550hPa         []float32 `json:"geopotential_height_550hPa"`
	GeopotentialHeight600hPa         []float32 `json:"geopotential_height_600hPa"`
	GeopotentialHeight650hPa         []float32 `json:"geopotential_height_650hPa"`
	GeopotentialHeight700hPa         []float32 `json:"geopotential_height_700hPa"`
	GeopotentialHeight750hPa         []float32 `json:"geopotential_height_750hPa"`
	GeopotentialHeight800hPa         []float32 `json:"geopotential_height_800hPa"`
	GeopotentialHeight850hPa         []float32 `json:"geopotential_height_850hPa"`
	GeopotentialHeight900hPa         []float32 `json:"geopotential_height_900hPa"`
	GeopotentialHeight925hPa         []float32 `json:"geopotential_height_925hPa"`
	GeopotentialHeight950hPa         []float32 `json:"geopotential_height_950hPa"`
	GlobalTiltedIrradiance           []float32 `json:"global_tilted_irradiance"`
	GlobalTiltedIrradianceInstant    []float32 `json:"global_tilted_irradiance_instant"`
	IcePelletsProbability            []float32 `json:"ice_pellets_probability"`
	IsDay                            []int     `json:"is_day"`
	KIndex                           []float32 `json:"k_index"`
	LiftedIndex                      []float32 `json:"lifted_index"`
	LightningDensity                 []float32 `json:"lightning_density"`
	LightningPotential               []float32 `json:"lightning_potential"`
	MassDensity8m                    []float32 `json:"mass_density_8m"`
	OceanCurrentDirection            []float32 `json:"ocean_current_direction"`
	OceanCurrentVelocity             []float32 `json:"ocean_current_velocity"`
	Precipitation                    []float32 `json:"precipitation"`
	PrecipitationProbability         []float32 `json:"precipitation_probability"`
	PrecipitationType                []int     `json:"precipitation_type"`
	PressureMsl                      []float32 `json:"pressure_msl"`
	Rain                             []float32 `json:"rain"`
	RainProbability                  []float32 `json:"rain_probability"`
	RelativeHumidity1000hPa          []float32 `json:"relative_humidity_1000hPa"`
	RelativeHumidity100hPa           []float32 `json:"relative_humidity_100hPa"`
	RelativeHumidity10hPa            []float32 `json:"relative_humidity_10hPa"`
	RelativeHumidity125hPa           []float32 `json:"relative_humidity_125hPa"`
	RelativeHumidity150hPa           []float32 `json:"relative_humidity_150hPa"`
	RelativeHumidity175hPa           []float32 `json:"relative_humidity_175hPa"`
	RelativeHumidity200hPa           []float32 `json:"relative_humidity_200hPa"`
	RelativeHumidity225hPa           []float32 `json:"relative_humidity_225hPa"`
	RelativeHumidity250hPa           []float32 `json:"relative_humidity_250hPa"`
	RelativeHumidity275hPa           []float32 `json:"relative_humidity_275hPa"`
	RelativeHumidity2m               []float32 `json:"relative_humidity_2m"`
	RelativeHumidity300hPa           []float32 `json:"relative_humidity_300hPa"`
	RelativeHumidity350hPa           []float32 `json:"relative_humidity_350hPa"`
	RelativeHumidity400hPa           []float32 `json:"relative_humidity_400hPa"`
	RelativeHumidity450hPa           []float32 `json:"relative_humidity_450hPa"`
	RelativeHumidity500hPa           []float32 `json:"relative_humidity_500hPa"`
	RelativeHumidity50hPa            []float32 `json:"relative_humidity_50hPa"`
	RelativeHumidity550hPa           []float32 `json:"relative_humidity_550hPa"`
	RelativeHumidity600hPa           []float32 `json:"relative_humidity_600hPa"`
	RelativeHumidity650hPa           []float32 `json:"relative_humidity_650hPa"`
	RelativeHumidity700hPa           []float32 `json:"relative_humidity_700hPa"`
	RelativeHumidity750hPa           []float32 `json:"relative_humidity_750hPa"`
	RelativeHumidity800hPa           []float32 `json:"relative_humidity_800hPa"`
	RelativeHumidity850hPa           []float32 `json:"relative_humidity_850hPa"`
	RelativeHumidity900hPa           []float32 `json:"relative_humidity_900hPa"`
	RelativeHumidity925hPa           []float32 `json:"relative_humidity_925hPa"`
	RelativeHumidity950hPa           []float32 `json:"relative_humidity_950hPa"`
	RoughnessLength                  []float32 `json:"roughness_length"`
	Runoff                           []float32 `json:"runoff"`
	SeaIceThickness                  []float32 `json:"sea_ice_thickness"`
	SeaLevelHeightMsl                []float32 `json:"sea_level_height_msl"`
	SeaSurfaceTemperature            []float32 `json:"sea_surface_temperature"`
	ShortwaveRadiation               []float32 `json:"shortwave_radiation"`
	ShortwaveRadiationInstant        []float32 `json:"shortwave_radiation_instant"`
	Showers                          []float32 `json:"showers"`
	SnowDepth                        []float32 `json:"snow_depth"`
	SnowDepthWaterEquivalent         []float32 `json:"snow_depth_water_equivalent"`
	SnowHeight                       []float32 `json:"snow_height"`
	Snowfall                         []float32 `json:"snowfall"`
	SnowfallHeight                   []float32 `json:"snowfall_height"`
	SnowfallProbability              []float32 `json:"snowfall_probability"`
	SnowfallWaterEquivalent          []float32 `json:"snowfall_water_equivalent"`
	SoilMoisture0To10cm              []float32 `json:"soil_moisture_0_to_10cm"`
	SoilMoisture0To1cm               []float32 `json:"soil_moisture_0_to_1cm"`
	SoilMoisture0To7cm               []float32 `json:"soil_moisture_0_to_7cm"`
	SoilMoisture100To200cm           []float32 `json:"soil_moisture_100_to_200cm"`
	SoilMoisture100To255cm           []float32 `json:"soil_moisture_100_to_255cm"`
	SoilMoisture100To300cm           []float32 `json:"soil_moisture_100_to_300cm"`
	SoilMoisture10To35cm             []float32 `json:"soil_moisture_10_to_35cm"`
	SoilMoisture10To40cm             []float32 `json:"soil_moisture_10_to_40cm"`
	SoilMoisture1To3cm               []float32 `json:"soil_moisture_1_to_3cm"`
	SoilMoisture243To729cm           []float32 `json:"soil_moisture_243_to_729cm"`
	SoilMoisture27To81cm             []float32 `json:"soil_moisture_27_to_81cm"`
	SoilMoisture28To100cm            []float32 `json:"soil_moisture_28_to_100cm"`
	SoilMoisture35To100cm            []float32 `json:"soil_moisture_35_to_100cm"`
	SoilMoisture3To9cm               []float32 `json:"soil_moisture_3_to_9cm"`
	SoilMoisture40To100cm            []float32 `json:"soil_moisture_40_to_100cm"`
	SoilMoisture729To2187cm          []float32 `json:"soil_moisture_729_to_2187cm"`
	SoilMoisture7To28cm              []float32 `json:"soil_moisture_7_to_28cm"`
	SoilMoisture81To243cm            []float32 `json:"soil_moisture_81_to_243cm"`
	SoilMoisture9To27cm              []float32 `json:"soil_moisture_9_to_27cm"`
	SoilTemperature0To10cm           []float32 `json:"soil_temperature_0_to_10cm"`
	SoilTemperature0To7cm            []float32 `json:"soil_temperature_0_to_7cm"`
	SoilTemperature0cm               []float32 `json:"soil_temperature_0cm"`
	SoilTemperature100To200cm        []float32 `json:"soil_temperature_100_to_200cm"`
	SoilTemperature100To255cm        []float32 `json:"soil_temperature_100_to_255cm"`
	SoilTemperature100To300cm        []float32 `json:"soil_temperature_100_to_300cm"`
	SoilTemperature10To35cm          []float32 `json:"soil_temperature_10_to_35cm"`
	SoilTemperature10To40cm          []float32 `json:"soil_temperature_10_to_40cm"`
	SoilTemperature1458cm            []float32 `json:"soil_temperature_1458cm"`
	SoilTemperature162cm             []float32 `json:"soil_temperature_162cm"`
	SoilTemperature18cm              []float32 `json:"soil_temperature_18cm"`
	SoilTemperature28To100cm         []float32 `json:"soil_temperature_28_to_100cm"`
	SoilTemperature35To100cm         []float32 `json:"soil_temperature_35_to_100cm"`
	SoilTemperature40To100cm         []float32 `json:"soil_temperature_40_to_100cm"`
	SoilTemperature486cm             []float32 `json:"soil_temperature_486cm"`
	SoilTemperature54cm              []float32 `json:"soil_temperature_54cm"`
	SoilTemperature6cm               []float32 `json:"soil_temperature_6cm"`
	SoilTemperature7To28cm           []float32 `json:"soil_temperature_7_to_28cm"`
	SunshineDuration                 []float32 `json:"sunshine_duration"`
	SurfacePressure                  []float32 `json:"surface_pressure"`
	SurfaceTemperature               []float32 `json:"surface_temperature"`
	Temperature1000hPa               []float32 `json:"temperature_1000hPa"`
	Temperature100hPa                []float32 `json:"temperature_100hPa"`
	Temperature100m                  []float32 `json:"temperature_100m"`
	Temperature10hPa                 []float32 `json:"temperature_10hPa"`
	Temperature120m                  []float32 `json:"temperature_120m"`
	Temperature125hPa                []float32 `json:"temperature_125hPa"`
	Temperature150hPa                []float32 `json:"temperature_150hPa"`
	Temperature150m                  []float32 `json:"temperature_150m"`
	Temperature175hPa                []float32 `json:"temperature_175hPa"`
	Temperature180m                  []float32 `json:"temperature_180m"`
	Temperature200hPa                []float32 `json:"temperature_200hPa"`
	Temperature200m                  []float32 `json:"temperature_200m"`
	Temperature20m                   []float32 `json:"temperature_20m"`
	Temperature225hPa                []float32 `json:"temperature_225hPa"`
	Temperature250hPa                []float32 `json:"temperature_250hPa"`
	Temperature275hPa                []float32 `json:"temperature_275hPa"`
	Temperature2m                    []float32 `json:"temperature_2m"`
	Temperature2mMax                 []float32 `json:"temperature_2m_max"`
	Temperature2mMin                 []float32 `json:"temperature_2m_min"`
	Temperature300hPa                []float32 `json:"temperature_300hPa"`
	Temperature350hPa                []float32 `json:"temperature_350hPa"`
	Temperature400hPa                []float32 `json:"temperature_400hPa"`
	Temperature40m                   []float32 `json:"temperature_40m"`
	Temperature450hPa                []float32 `json:"temperature_450hPa"`
	Temperature500hPa                []float32 `json:"temperature_500hPa"`
	Temperature50hPa                 []float32 `json:"temperature_50hPa"`
	Temperature50m                   []float32 `json:"temperature_50m"`
	Temperature550hPa                []float32 `json:"temperature_550hPa"`
	Temperature600hPa                []float32 `json:"temperature_600hPa"`
	Temperature650hPa                []float32 `json:"temperature_650hPa"`
	Temperature700hPa                []float32 `json:"temperature_700hPa"`
	Temperature750hPa                []float32 `json:"temperature_750hPa"`
	Temperature800hPa                []float32 `json:"temperature_800hPa"`
	Temperature80m                   []float32 `json:"temperature_80m"`
	Temperature850hPa                []float32 `json:"temperature_850hPa"`
	Temperature900hPa                []float32 `json:"temperature_900hPa"`
	Temperature925hPa                []float32 `json:"temperature_925hPa"`
	Temperature950hPa                []float32 `json:"temperature_950hPa"`
	TerrestrialRadiation             []float32 `json:"terrestrial_radiation"`
	TerrestrialRadiationInstant      []float32 `json:"terrestrial_radiation_instant"`
	ThunderstormProbability          []float32 `json:"thunderstorm_probability"`
	Time                             []string  `json:"time"`
	TotalColumnIntegratedWaterVapour []float32 `json:"total_column_integrated_water_vapour"`
	Updraft                          []float32 `json:"updraft"`
	UvIndex                          []float32 `json:"uv_index"`
	UvIndexClearSky                  []float32 `json:"uv_index_clear_sky"`
	VapourPressureDeficit            []float32 `json:"vapour_pressure_deficit"`
	VerticalVelocity1000hPa          []float32 `json:"vertical_velocity_1000hPa"`
	VerticalVelocity100hPa           []float32 `json:"vertical_velocity_100hPa"`
	VerticalVelocity10hPa            []float32 `json:"vertical_velocity_10hPa"`
	VerticalVelocity125hPa           []float32 `json:"vertical_velocity_125hPa"`
	VerticalVelocity150hPa           []float32 `json:"vertical_velocity_150hPa"`
	VerticalVelocity175hPa           []float32 `json:"vertical_velocity_175hPa"`
	VerticalVelocity200hPa           []float32 `json:"vertical_velocity_200hPa"`
	VerticalVelocity225hPa           []float32 `json:"vertical_velocity_225hPa"`
	VerticalVelocity250hPa           []float32 `json:"vertical_velocity_250hPa"`
	VerticalVelocity275hPa           []float32 `json:"vertical_velocity_275hPa"`
	VerticalVelocity300hPa           []float32 `json:"vertical_velocity_300hPa"`
	VerticalVelocity350hPa           []float32 `json:"vertical_velocity_350hPa"`
	VerticalVelocity400hPa           []float32 `json:"vertical_velocity_400hPa"`
	VerticalVelocity450hPa           []float32 `json:"vertical_velocity_450hPa"`
	VerticalVelocity500hPa           []float32 `json:"vertical_velocity_500hPa"`
	VerticalVelocity50hPa            []float32 `json:"vertical_velocity_50hPa"`
	VerticalVelocity550hPa           []float32 `json:"vertical_velocity_550hPa"`
	VerticalVelocity600hPa           []float32 `json:"vertical_velocity_600hPa"`
	VerticalVelocity650hPa           []float32 `json:"vertical_velocity_650hPa"`
	VerticalVelocity700hPa           []float32 `json:"vertical_velocity_700hPa"`
	VerticalVelocity750hPa           []float32 `json:"vertical_velocity_750hPa"`
	VerticalVelocity800hPa           []float32 `json:"vertical_velocity_800hPa"`
	VerticalVelocity850hPa           []float32 `json:"vertical_velocity_850hPa"`
	VerticalVelocity900hPa           []float32 `json:"vertical_velocity_900hPa"`
	VerticalVelocity925hPa           []float32 `json:"vertical_velocity_925hPa"`
	VerticalVelocity950hPa           []float32 `json:"vertical_velocity_950hPa"`
	Visibility                       []float32 `json:"visibility"`
	WeatherCode                      []int     `json:"weather_code"`
	WetBulbTemperature2m             []float32 `json:"wet_bulb_temperature_2m"`
	WindDirection1000hPa             []float32 `json:"wind_direction_1000hPa"`
	WindDirection100hPa              []float32 `json:"wind_direction_100hPa"`
	WindDirection100m                []float32 `json:"wind_direction_100m"`
	WindDirection10hPa               []float32 `json:"wind_direction_10hPa"`
	WindDirection10m                 []float32 `json:"wind_direction_10m"`
	WindDirection120m                []float32 `json:"wind_direction_120m"`
	WindDirection125hPa              []float32 `json:"wind_direction_125hPa"`
	WindDirection140m                []float32 `json:"wind_direction_140m"`
	WindDirection150hPa              []float32 `json:"wind_direction_150hPa"`
	WindDirection150m                []float32 `json:"wind_direction_150m"`
	WindDirection160m                []float32 `json:"wind_direction_160m"`
	WindDirection175hPa              []float32 `json:"wind_direction_175hPa"`
	WindDirection180m                []float32 `json:"wind_direction_180m"`
	WindDirection200hPa              []float32 `json:"wind_direction_200hPa"`
	WindDirection200m                []float32 `json:"wind_direction_200m"`
	WindDirection20m                 []float32 `json:"wind_direction_20m"`
	WindDirection225hPa              []float32 `json:"wind_direction_225hPa"`
	WindDirection250hPa              []float32 `json:"wind_direction_250hPa"`
	WindDirection275hPa              []float32 `json:"wind_direction_275hPa"`
	WindDirection300hPa              []float32 `json:"wind_direction_300hPa"`
	WindDirection30m                 []float32 `json:"wind_direction_30m"`
	WindDirection350hPa              []float32 `json:"wind_direction_350hPa"`
	WindDirection400hPa              []float32 `json:"wind_direction_400hPa"`
	WindDirection40m                 []float32 `json:"wind_direction_40m"`
	WindDirection450hPa              []float32 `json:"wind_direction_450hPa"`
	WindDirection500hPa              []float32 `json:"wind_direction_500hPa"`
	WindDirection50hPa               []float32 `json:"wind_direction_50hPa"`
	WindDirection50m                 []float32 `json:"wind_direction_50m"`
	WindDirection550hPa              []float32 `json:"wind_direction_550hPa"`
	WindDirection600hPa              []float32 `json:"wind_direction_600hPa"`
	WindDirection650hPa              []float32 `json:"wind_direction_650hPa"`
	WindDirection700hPa              []float32 `json:"wind_direction_700hPa"`
	WindDirection70m                 []float32 `json:"wind_direction_70m"`
	WindDirection750hPa              []float32 `json:"wind_direction_750hPa"`
	WindDirection800hPa              []float32 `json:"wind_direction_800hPa"`
	WindDirection80m                 []float32 `json:"wind_direction_80m"`
	WindDirection850hPa              []float32 `json:"wind_direction_850hPa"`
	WindDirection900hPa              []float32 `json:"wind_direction_900hPa"`
	WindDirection925hPa              []float32 `json:"wind_direction_925hPa"`
	WindDirection950hPa              []float32 `json:"wind_direction_950hPa"`
	WindGusts10m                     []float32 `json:"wind_gusts_10m"`
	WindSpeed1000hPa                 []float32 `json:"wind_speed_1000hPa"`
	WindSpeed100hPa                  []float32 `json:"wind_speed_100hPa"`
	WindSpeed100m                    []float32 `json:"wind_speed_100m"`
	WindSpeed10hPa                   []float32 `json:"wind_speed_10hPa"`
	WindSpeed10m                     []float32 `json:"wind_speed_10m"`
	WindSpeed120m                    []float32 `json:"wind_speed_120m"`
	WindSpeed125hPa                  []float32 `json:"wind_speed_125hPa"`
	WindSpeed140m                    []float32 `json:"wind_speed_140m"`
	WindSpeed150hPa                  []float32 `json:"wind_speed_150hPa"`
	WindSpeed150m                    []float32 `json:"wind_speed_150m"`
	WindSpeed160m                    []float32 `json:"wind_speed_160m"`
	WindSpeed175hPa                  []float32 `json:"wind_speed_175hPa"`
	WindSpeed180m                    []float32 `json:"wind_speed_180m"`
	WindSpeed200hPa                  []float32 `json:"wind_speed_200hPa"`
	WindSpeed200m                    []float32 `json:"wind_speed_200m"`
	WindSpeed20m                     []float32 `json:"wind_speed_20m"`
	WindSpeed225hPa                  []float32 `json:"wind_speed_225hPa"`
	WindSpeed250hPa                  []float32 `json:"wind_speed_250hPa"`
	WindSpeed275hPa                  []float32 `json:"wind_speed_275hPa"`
	WindSpeed300hPa                  []float32 `json:"wind_speed_300hPa"`
	WindSpeed30m                     []float32 `json:"wind_speed_30m"`
	WindSpeed350hPa                  []float32 `json:"wind_speed_350hPa"`
	WindSpeed400hPa                  []float32 `json:"wind_speed_400hPa"`
	WindSpeed40m                     []float32 `json:"wind_speed_40m"`
	WindSpeed450hPa                  []float32 `json:"wind_speed_450hPa"`
	WindSpeed500hPa                  []float32 `json:"wind_speed_500hPa"`
	WindSpeed50hPa                   []float32 `json:"wind_speed_50hPa"`
	WindSpeed50m                     []float32 `json:"wind_speed_50m"`
	WindSpeed550hPa                  []float32 `json:"wind_speed_550hPa"`
	WindSpeed600hPa                  []float32 `json:"wind_speed_600hPa"`
	WindSpeed650hPa                  []float32 `json:"wind_speed_650hPa"`
	WindSpeed700hPa                  []float32 `json:"wind_speed_700hPa"`
	WindSpeed70m                     []float32 `json:"wind_speed_70m"`
	WindSpeed750hPa                  []float32 `json:"wind_speed_750hPa"`
	WindSpeed800hPa                  []float32 `json:"wind_speed_800hPa"`
	WindSpeed80m                     []float32 `json:"wind_speed_80m"`
	WindSpeed850hPa                  []float32 `json:"wind_speed_850hPa"`
	WindSpeed900hPa                  []float32 `json:"wind_speed_900hPa"`
	WindSpeed925hPa                  []float32 `json:"wind_speed_925hPa"`
	WindSpeed950hPa                  []float32 `json:"wind_speed_950hPa"`
}

type HourlyUnits struct {
	Albedo                           string `json:"albedo"`
	ApparentTemperature              string `json:"apparent_temperature"`
	BoundaryLayerHeight              string `json:"boundary_layer_height"`
	Cape                             string `json:"cape"`
	CloudCover                       string `json:"cloud_cover"`
	CloudCover1000hPa                string `json:"cloud_cover_1000hPa"`
	CloudCover100hPa                 string `json:"cloud_cover_100hPa"`
	CloudCover10hPa                  string `json:"cloud_cover_10hPa"`
	CloudCover125hPa                 string `json:"cloud_cover_125hPa"`
	CloudCover150hPa                 string `json:"cloud_cover_150hPa"`
	CloudCover175hPa                 string `json:"cloud_cover_175hPa"`
	CloudCover200hPa                 string `json:"cloud_cover_200hPa"`
	CloudCover225hPa                 string `json:"cloud_cover_225hPa"`
	CloudCover250hPa                 string `json:"cloud_cover_250hPa"`
	CloudCover275hPa                 string `json:"cloud_cover_275hPa"`
	CloudCover300hPa                 string `json:"cloud_cover_300hPa"`
	CloudCover350hPa                 string `json:"cloud_cover_350hPa"`
	CloudCover400hPa                 string `json:"cloud_cover_400hPa"`
	CloudCover450hPa                 string `json:"cloud_cover_450hPa"`
	CloudCover500hPa                 string `json:"cloud_cover_500hPa"`
	CloudCover50hPa                  string `json:"cloud_cover_50hPa"`
	CloudCover550hPa                 string `json:"cloud_cover_550hPa"`
	CloudCover600hPa                 string `json:"cloud_cover_600hPa"`
	CloudCover650hPa                 string `json:"cloud_cover_650hPa"`
	CloudCover700hPa                 string `json:"cloud_cover_700hPa"`
	CloudCover750hPa                 string `json:"cloud_cover_750hPa"`
	CloudCover800hPa                 string `json:"cloud_cover_800hPa"`
	CloudCover850hPa                 string `json:"cloud_cover_850hPa"`
	CloudCover900hPa                 string `json:"cloud_cover_900hPa"`
	CloudCover925hPa                 string `json:"cloud_cover_925hPa"`
	CloudCover950hPa                 string `json:"cloud_cover_950hPa"`
	CloudCoverHigh                   string `json:"cloud_cover_high"`
	CloudCoverLow                    string `json:"cloud_cover_low"`
	CloudCoverMid                    string `json:"cloud_cover_mid"`
	ConvectiveCloudBase              string `json:"convective_cloud_base"`
	ConvectiveCloudTop               string `json:"convective_cloud_top"`
	ConvectiveInhibition             string `json:"convective_inhibition"`
	DewPoint1000hPa                  string `json:"dew_point_1000hPa"`
	DewPoint100hPa                   string `json:"dew_point_100hPa"`
	DewPoint10hPa                    string `json:"dew_point_10hPa"`
	DewPoint125hPa                   string `json:"dew_point_125hPa"`
	DewPoint150hPa                   string `json:"dew_point_150hPa"`
	DewPoint175hPa                   string `json:"dew_point_175hPa"`
	DewPoint200hPa                   string `json:"dew_point_200hPa"`
	DewPoint225hPa                   string `json:"dew_point_225hPa"`
	DewPoint250hPa                   string `json:"dew_point_250hPa"`
	DewPoint275hPa                   string `json:"dew_point_275hPa"`
	DewPoint2m                       string `json:"dew_point_2m"`
	DewPoint300hPa                   string `json:"dew_point_300hPa"`
	DewPoint350hPa                   string `json:"dew_point_350hPa"`
	DewPoint400hPa                   string `json:"dew_point_400hPa"`
	DewPoint450hPa                   string `json:"dew_point_450hPa"`
	DewPoint500hPa                   string `json:"dew_point_500hPa"`
	DewPoint50hPa                    string `json:"dew_point_50hPa"`
	DewPoint550hPa                   string `json:"dew_point_550hPa"`
	DewPoint600hPa                   string `json:"dew_point_600hPa"`
	DewPoint650hPa                   string `json:"dew_point_650hPa"`
	DewPoint700hPa                   string `json:"dew_point_700hPa"`
	DewPoint750hPa                   string `json:"dew_point_750hPa"`
	DewPoint800hPa                   string `json:"dew_point_800hPa"`
	DewPoint850hPa                   string `json:"dew_point_850hPa"`
	DewPoint900hPa                   string `json:"dew_point_900hPa"`
	DewPoint925hPa                   string `json:"dew_point_925hPa"`
	DewPoint950hPa                   string `json:"dew_point_950hPa"`
	DiffuseRadiation                 string `json:"diffuse_radiation"`
	DiffuseRadiationInstant          string `json:"diffuse_radiation_instant"`
	DirectNormalIrradiance           string `json:"direct_normal_irradiance"`
	DirectNormalIrradianceInstant    string `json:"direct_normal_irradiance_instant"`
	DirectRadiation                  string `json:"direct_radiation"`
	DirectRadiationInstant           string `json:"direct_radiation_instant"`
	Et0FaoEvapotranspiration         string `json:"et0_fao_evapotranspiration"`
	Evapotranspiration               string `json:"evapotranspiration"`
	FreezingLevelHeight              string `json:"freezing_level_height"`
	FreezingRainProbability          string `json:"freezing_rain_probability"`
	GeopotentialHeight1000hPa        string `json:"geopotential_height_1000hPa"`
	GeopotentialHeight100hPa         string `json:"geopotential_height_100hPa"`
	GeopotentialHeight10hPa          string `json:"geopotential_height_10hPa"`
	GeopotentialHeight125hPa         string `json:"geopotential_height_125hPa"`
	GeopotentialHeight150hPa         string `json:"geopotential_height_150hPa"`
	GeopotentialHeight175hPa         string `json:"geopotential_height_175hPa"`
	GeopotentialHeight200hPa         string `json:"geopotential_height_200hPa"`
	GeopotentialHeight225hPa         string `json:"geopotential_height_225hPa"`
	GeopotentialHeight250hPa         string `json:"geopotential_height_250hPa"`
	GeopotentialHeight275hPa         string `json:"geopotential_height_275hPa"`
	GeopotentialHeight300hPa         string `json:"geopotential_height_300hPa"`
	GeopotentialHeight350hPa         string `json:"geopotential_height_350hPa"`
	GeopotentialHeight400hPa         string `json:"geopotential_height_400hPa"`
	GeopotentialHeight450hPa         string `json:"geopotential_height_450hPa"`
	GeopotentialHeight500hPa         string `json:"geopotential_height_500hPa"`
	GeopotentialHeight50hPa          string `json:"geopotential_height_50hPa"`
	GeopotentialHeight550hPa         string `json:"geopotential_height_550hPa"`
	GeopotentialHeight600hPa         string `json:"geopotential_height_600hPa"`
	GeopotentialHeight650hPa         string `json:"geopotential_height_650hPa"`
	GeopotentialHeight700hPa         string `json:"geopotential_height_700hPa"`
	GeopotentialHeight750hPa         string `json:"geopotential_height_750hPa"`
	GeopotentialHeight800hPa         string `json:"geopotential_height_800hPa"`
	GeopotentialHeight850hPa         string `json:"geopotential_height_850hPa"`
	GeopotentialHeight900hPa         string `json:"geopotential_height_900hPa"`
	GeopotentialHeight925hPa         string `json:"geopotential_height_925hPa"`
	GeopotentialHeight950hPa         string `json:"geopotential_height_950hPa"`
	GlobalTiltedIrradiance           string `json:"global_tilted_irradiance"`
	GlobalTiltedIrradianceInstant    string `json:"global_tilted_irradiance_instant"`
	IcePelletsProbability            string `json:"ice_pellets_probability"`
	IsDay                            string `json:"is_day"`
	KIndex                           string `json:"k_index"`
	LiftedIndex                      string `json:"lifted_index"`
	LightningDensity                 string `json:"lightning_density"`
	LightningPotential               string `json:"lightning_potential"`
	MassDensity8m                    string `json:"mass_density_8m"`
	OceanCurrentDirection            string `json:"ocean_current_direction"`
	OceanCurrentVelocity             string `json:"ocean_current_velocity"`
	Precipitation                    string `json:"precipitation"`
	PrecipitationProbability         string `json:"precipitation_probability"`
	PrecipitationType                string `json:"precipitation_type"`
	PressureMsl                      string `json:"pressure_msl"`
	Rain                             string `json:"rain"`
	RainProbability                  string `json:"rain_probability"`
	RelativeHumidity1000hPa          string `json:"relative_humidity_1000hPa"`
	RelativeHumidity100hPa           string `json:"relative_humidity_100hPa"`
	RelativeHumidity10hPa            string `json:"relative_humidity_10hPa"`
	RelativeHumidity125hPa           string `json:"relative_humidity_125hPa"`
	RelativeHumidity150hPa           string `json:"relative_humidity_150hPa"`
	RelativeHumidity175hPa           string `json:"relative_humidity_175hPa"`
	RelativeHumidity200hPa           string `json:"relative_humidity_200hPa"`
	RelativeHumidity225hPa           string `json:"relative_humidity_225hPa"`
	RelativeHumidity250hPa           string `json:"relative_humidity_250hPa"`
	RelativeHumidity275hPa           string `json:"relative_humidity_275hPa"`
	RelativeHumidity2m               string `json:"relative_humidity_2m"`
	RelativeHumidity300hPa           string `json:"relative_humidity_300hPa"`
	RelativeHumidity350hPa           string `json:"relative_humidity_350hPa"`
	RelativeHumidity400hPa           string `json:"relative_humidity_400hPa"`
	RelativeHumidity450hPa           string `json:"relative_humidity_450hPa"`
	RelativeHumidity500hPa           string `json:"relative_humidity_500hPa"`
	RelativeHumidity50hPa            string `json:"relative_humidity_50hPa"`
	RelativeHumidity550hPa           string `json:"relative_humidity_550hPa"`
	RelativeHumidity600hPa           string `json:"relative_humidity_600hPa"`
	RelativeHumidity650hPa           string `json:"relative_humidity_650hPa"`
	RelativeHumidity700hPa           string `json:"relative_humidity_700hPa"`
	RelativeHumidity750hPa           string `json:"relative_humidity_750hPa"`
	RelativeHumidity800hPa           string `json:"relative_humidity_800hPa"`
	RelativeHumidity850hPa           string `json:"relative_humidity_850hPa"`
	RelativeHumidity900hPa           string `json:"relative_humidity_900hPa"`
	RelativeHumidity925hPa           string `json:"relative_humidity_925hPa"`
	RelativeHumidity950hPa           string `json:"relative_humidity_950hPa"`
	RoughnessLength                  string `json:"roughness_length"`
	Runoff                           string `json:"runoff"`
	SeaIceThickness                  string `json:"sea_ice_thickness"`
	SeaLevelHeightMsl                string `json:"sea_level_height_msl"`
	SeaSurfaceTemperature            string `json:"sea_surface_temperature"`
	ShortwaveRadiation               string `json:"shortwave_radiation"`
	ShortwaveRadiationInstant        string `json:"shortwave_radiation_instant"`
	Showers                          string `json:"showers"`
	SnowDepth                        string `json:"snow_depth"`
	SnowDepthWaterEquivalent         string `json:"snow_depth_water_equivalent"`
	SnowHeight                       string `json:"snow_height"`
	Snowfall                         string `json:"snowfall"`
	SnowfallHeight                   string `json:"snowfall_height"`
	SnowfallProbability              string `json:"snowfall_probability"`
	SnowfallWaterEquivalent          string `json:"snowfall_water_equivalent"`
	SoilMoisture0To10cm              string `json:"soil_moisture_0_to_10cm"`
	SoilMoisture0To1cm               string `json:"soil_moisture_0_to_1cm"`
	SoilMoisture0To7cm               string `json:"soil_moisture_0_to_7cm"`
	SoilMoisture100To200cm           string `json:"soil_moisture_100_to_200cm"`
	SoilMoisture100To255cm           string `json:"soil_moisture_100_to_255cm"`
	SoilMoisture100To300cm           string `json:"soil_moisture_100_to_300cm"`
	SoilMoisture10To35cm             string `json:"soil_moisture_10_to_35cm"`
	SoilMoisture10To40cm             string `json:"soil_moisture_10_to_40cm"`
	SoilMoisture1To3cm               string `json:"soil_moisture_1_to_3cm"`
	SoilMoisture243To729cm           string `json:"soil_moisture_243_to_729cm"`
	SoilMoisture27To81cm             string `json:"soil_moisture_27_to_81cm"`
	SoilMoisture28To100cm            string `json:"soil_moisture_28_to_100cm"`
	SoilMoisture35To100cm            string `json:"soil_moisture_35_to_100cm"`
	SoilMoisture3To9cm               string `json:"soil_moisture_3_to_9cm"`
	SoilMoisture40To100cm            string `json:"soil_moisture_40_to_100cm"`
	SoilMoisture729To2187cm          string `json:"soil_moisture_729_to_2187cm"`
	SoilMoisture7To28cm              string `json:"soil_moisture_7_to_28cm"`
	SoilMoisture81To243cm            string `json:"soil_moisture_81_to_243cm"`
	SoilMoisture9To27cm              string `json:"soil_moisture_9_to_27cm"`
	SoilTemperature0To10cm           string `json:"soil_temperature_0_to_10cm"`
	SoilTemperature0To7cm            string `json:"soil_temperature_0_to_7cm"`
	SoilTemperature0cm               string `json:"soil_temperature_0cm"`
	SoilTemperature100To200cm        string `json:"soil_temperature_100_to_200cm"`
	SoilTemperature100To255cm        string `json:"soil_temperature_100_to_255cm"`
	SoilTemperature100To300cm        string `json:"soil_temperature_100_to_300cm"`
	SoilTemperature10To35cm          string `json:"soil_temperature_10_to_35cm"`
	SoilTemperature10To40cm          string `json:"soil_temperature_10_to_40cm"`
	SoilTemperature1458cm            string `json:"soil_temperature_1458cm"`
	SoilTemperature162cm             string `json:"soil_temperature_162cm"`
	SoilTemperature18cm              string `json:"soil_temperature_18cm"`
	SoilTemperature28To100cm         string `json:"soil_temperature_28_to_100cm"`
	SoilTemperature35To100cm         string `json:"soil_temperature_35_to_100cm"`
	SoilTemperature40To100cm         string `json:"soil_temperature_40_to_100cm"`
	SoilTemperature486cm             string `json:"soil_temperature_486cm"`
	SoilTemperature54cm              string `json:"soil_temperature_54cm"`
	SoilTemperature6cm               string `json:"soil_temperature_6cm"`
	SoilTemperature7To28cm           string `json:"soil_temperature_7_to_28cm"`
	SunshineDuration                 string `json:"sunshine_duration"`
	SurfacePressure                  string `json:"surface_pressure"`
	SurfaceTemperature               string `json:"surface_temperature"`
	Temperature1000hPa               string `json:"temperature_1000hPa"`
	Temperature100hPa                string `json:"temperature_100hPa"`
	Temperature100m                  string `json:"temperature_100m"`
	Temperature10hPa                 string `json:"temperature_10hPa"`
	Temperature120m                  string `json:"temperature_120m"`
	Temperature125hPa                string `json:"temperature_125hPa"`
	Temperature150hPa                string `json:"temperature_150hPa"`
	Temperature150m                  string `json:"temperature_150m"`
	Temperature175hPa                string `json:"temperature_175hPa"`
	Temperature180m                  string `json:"temperature_180m"`
	Temperature200hPa                string `json:"temperature_200hPa"`
	Temperature200m                  string `json:"temperature_200m"`
	Temperature20m                   string `json:"temperature_20m"`
	Temperature225hPa                string `json:"temperature_225hPa"`
	Temperature250hPa                string `json:"temperature_250hPa"`
	Temperature275hPa                string `json:"temperature_275hPa"`
	Temperature2m                    string `json:"temperature_2m"`
	Temperature2mMax                 string `json:"temperature_2m_max"`
	Temperature2mMin                 string `json:"temperature_2m_min"`
	Temperature300hPa                string `json:"temperature_300hPa"`
	Temperature350hPa                string `json:"temperature_350hPa"`
	Temperature400hPa                string `json:"temperature_400hPa"`
	Temperature40m                   string `json:"temperature_40m"`
	Temperature450hPa                string `json:"temperature_450hPa"`
	Temperature500hPa                string `json:"temperature_500hPa"`
	Temperature50hPa                 string `json:"temperature_50hPa"`
	Temperature50m                   string `json:"temperature_50m"`
	Temperature550hPa                string `json:"temperature_550hPa"`
	Temperature600hPa                string `json:"temperature_600hPa"`
	Temperature650hPa                string `json:"temperature_650hPa"`
	Temperature700hPa                string `json:"temperature_700hPa"`
	Temperature750hPa                string `json:"temperature_750hPa"`
	Temperature800hPa                string `json:"temperature_800hPa"`
	Temperature80m                   string `json:"temperature_80m"`
	Temperature850hPa                string `json:"temperature_850hPa"`
	Temperature900hPa                string `json:"temperature_900hPa"`
	Temperature925hPa                string `json:"temperature_925hPa"`
	Temperature950hPa                string `json:"temperature_950hPa"`
	TerrestrialRadiation             string `json:"terrestrial_radiation"`
	TerrestrialRadiationInstant      string `json:"terrestrial_radiation_instant"`
	ThunderstormProbability          string `json:"thunderstorm_probability"`
	Time                             string `json:"time"`
	TotalColumnIntegratedWaterVapour string `json:"total_column_integrated_water_vapour"`
	Updraft                          string `json:"updraft"`
	UvIndex                          string `json:"uv_index"`
	UvIndexClearSky                  string `json:"uv_index_clear_sky"`
	VapourPressureDeficit            string `json:"vapour_pressure_deficit"`
	VerticalVelocity1000hPa          string `json:"vertical_velocity_1000hPa"`
	VerticalVelocity100hPa           string `json:"vertical_velocity_100hPa"`
	VerticalVelocity10hPa            string `json:"vertical_velocity_10hPa"`
	VerticalVelocity125hPa           string `json:"vertical_velocity_125hPa"`
	VerticalVelocity150hPa           string `json:"vertical_velocity_150hPa"`
	VerticalVelocity175hPa           string `json:"vertical_velocity_175hPa"`
	VerticalVelocity200hPa           string `json:"vertical_velocity_200hPa"`
	VerticalVelocity225hPa           string `json:"vertical_velocity_225hPa"`
	VerticalVelocity250hPa           string `json:"vertical_velocity_250hPa"`
	VerticalVelocity275hPa           string `json:"vertical_velocity_275hPa"`
	VerticalVelocity300hPa           string `json:"vertical_velocity_300hPa"`
	VerticalVelocity350hPa           string `json:"vertical_velocity_350hPa"`
	VerticalVelocity400hPa           string `json:"vertical_velocity_400hPa"`
	VerticalVelocity450hPa           string `json:"vertical_velocity_450hPa"`
	VerticalVelocity500hPa           string `json:"vertical_velocity_500hPa"`
	VerticalVelocity50hPa            string `json:"vertical_velocity_50hPa"`
	VerticalVelocity550hPa           string `json:"vertical_velocity_550hPa"`
	VerticalVelocity600hPa           string `json:"vertical_velocity_600hPa"`
	VerticalVelocity650hPa           string `json:"vertical_velocity_650hPa"`
	VerticalVelocity700hPa           string `json:"vertical_velocity_700hPa"`
	VerticalVelocity750hPa           string `json:"vertical_velocity_750hPa"`
	VerticalVelocity800hPa           string `json:"vertical_velocity_800hPa"`
	VerticalVelocity850hPa           string `json:"vertical_velocity_850hPa"`
	VerticalVelocity900hPa           string `json:"vertical_velocity_900hPa"`
	VerticalVelocity925hPa           string `json:"vertical_velocity_925hPa"`
	VerticalVelocity950hPa           string `json:"vertical_velocity_950hPa"`
	Visibility                       string `json:"visibility"`
	WeatherCode                      string `json:"weather_code"`
	WetBulbTemperature2m             string `json:"wet_bulb_temperature_2m"`
	WindDirection1000hPa             string `json:"wind_direction_1000hPa"`
	WindDirection100hPa              string `json:"wind_direction_100hPa"`
	WindDirection100m                string `json:"wind_direction_100m"`
	WindDirection10hPa               string `json:"wind_direction_10hPa"`
	WindDirection10m                 string `json:"wind_direction_10m"`
	WindDirection120m                string `json:"wind_direction_120m"`
	WindDirection125hPa              string `json:"wind_direction_125hPa"`
	WindDirection140m                string `json:"wind_direction_140m"`
	WindDirection150hPa              string `json:"wind_direction_150hPa"`
	WindDirection150m                string `json:"wind_direction_150m"`
	WindDirection160m                string `json:"wind_direction_160m"`
	WindDirection175hPa              string `json:"wind_direction_175hPa"`
	WindDirection180m                string `json:"wind_direction_180m"`
	WindDirection200hPa              string `json:"wind_direction_200hPa"`
	WindDirection200m                string `json:"wind_direction_200m"`
	WindDirection20m                 string `json:"wind_direction_20m"`
	WindDirection225hPa              string `json:"wind_direction_225hPa"`
	WindDirection250hPa              string `json:"wind_direction_250hPa"`
	WindDirection275hPa              string `json:"wind_direction_275hPa"`
	WindDirection300hPa              string `json:"wind_direction_300hPa"`
	WindDirection30m                 string `json:"wind_direction_30m"`
	WindDirection350hPa              string `json:"wind_direction_350hPa"`
	WindDirection400hPa              string `json:"wind_direction_400hPa"`
	WindDirection40m                 string `json:"wind_direction_40m"`
	WindDirection450hPa              string `json:"wind_direction_450hPa"`
	WindDirection500hPa              string `json:"wind_direction_500hPa"`
	WindDirection50hPa               string `json:"wind_direction_50hPa"`
	WindDirection50m                 string `json:"wind_direction_50m"`
	WindDirection550hPa              string `json:"wind_direction_550hPa"`
	WindDirection600hPa              string `json:"wind_direction_600hPa"`
	WindDirection650hPa              string `json:"wind_direction_650hPa"`
	WindDirection700hPa              string `json:"wind_direction_700hPa"`
	WindDirection70m                 string `json:"wind_direction_70m"`
	WindDirection750hPa              string `json:"wind_direction_750hPa"`
	WindDirection800hPa              string `json:"wind_direction_800hPa"`
	WindDirection80m                 string `json:"wind_direction_80m"`
	WindDirection850hPa              string `json:"wind_direction_850hPa"`
	WindDirection900hPa              string `json:"wind_direction_900hPa"`
	WindDirection925hPa              string `json:"wind_direction_925hPa"`
	WindDirection950hPa              string `json:"wind_direction_950hPa"`
	WindGusts10m                     string `json:"wind_gusts_10m"`
	WindSpeed1000hPa                 string `json:"wind_speed_1000hPa"`
	WindSpeed100hPa                  string `json:"wind_speed_100hPa"`
	WindSpeed100m                    string `json:"wind_speed_100m"`
	WindSpeed10hPa                   string `json:"wind_speed_10hPa"`
	WindSpeed10m                     string `json:"wind_speed_10m"`
	WindSpeed120m                    string `json:"wind_speed_120m"`
	WindSpeed125hPa                  string `json:"wind_speed_125hPa"`
	WindSpeed140m                    string `json:"wind_speed_140m"`
	WindSpeed150hPa                  string `json:"wind_speed_150hPa"`
	WindSpeed150m                    string `json:"wind_speed_150m"`
	WindSpeed160m                    string `json:"wind_speed_160m"`
	WindSpeed175hPa                  string `json:"wind_speed_175hPa"`
	WindSpeed180m                    string `json:"wind_speed_180m"`
	WindSpeed200hPa                  string `json:"wind_speed_200hPa"`
	WindSpeed200m                    string `json:"wind_speed_200m"`
	WindSpeed20m                     string `json:"wind_speed_20m"`
	WindSpeed225hPa                  string `json:"wind_speed_225hPa"`
	WindSpeed250hPa                  string `json:"wind_speed_250hPa"`
	WindSpeed275hPa                  string `json:"wind_speed_275hPa"`
	WindSpeed300hPa                  string `json:"wind_speed_300hPa"`
	WindSpeed30m                     string `json:"wind_speed_30m"`
	WindSpeed350hPa                  string `json:"wind_speed_350hPa"`
	WindSpeed400hPa                  string `json:"wind_speed_400hPa"`
	WindSpeed40m                     string `json:"wind_speed_40m"`
	WindSpeed450hPa                  string `json:"wind_speed_450hPa"`
	WindSpeed500hPa                  string `json:"wind_speed_500hPa"`
	WindSpeed50hPa                   string `json:"wind_speed_50hPa"`
	WindSpeed50m                     string `json:"wind_speed_50m"`
	WindSpeed550hPa                  string `json:"wind_speed_550hPa"`
	WindSpeed600hPa                  string `json:"wind_speed_600hPa"`
	WindSpeed650hPa                  string `json:"wind_speed_650hPa"`
	WindSpeed700hPa                  string `json:"wind_speed_700hPa"`
	WindSpeed70m                     string `json:"wind_speed_70m"`
	WindSpeed750hPa                  string `json:"wind_speed_750hPa"`
	WindSpeed800hPa                  string `json:"wind_speed_800hPa"`
	WindSpeed80m                     string `json:"wind_speed_80m"`
	WindSpeed850hPa                  string `json:"wind_speed_850hPa"`
	WindSpeed900hPa                  string `json:"wind_speed_900hPa"`
	WindSpeed925hPa                  string `json:"wind_speed_925hPa"`
	WindSpeed950hPa                  string `json:"wind_speed_950hPa"`
}

type Minutely15 struct {
	ApparentTemperature []float32 `json:"apparent_temperature"`
	Cape                []float32 `json:"cape"`
	DewPoint2m          []float32 `json:"dew_point_2m"`
	FreezingLevelHeight []float32 `json:"freezing_level_height"`
	IsDay               []int     `json:"is_day"`
	LightningPotential  []float32 `json:"lightning_potential"`
	Precipitation       []float32 `json:"precipitation"`
	Rain                []float32 `json:"rain"`
	RelativeHumidity2m  []float32 `json:"relative_humidity_2m"`
	Snowfall            []float32 `json:"snowfall"`
	SnowfallHeight      []float32 `json:"snowfall_height"`
	SunshineDuration    []float32 `json:"sunshine_duration"`
	Temperature2m       []float32 `json:"temperature_2m"`
	Time                []string  `json:"time"`
	Visibility          []float32 `json:"visibility"`
	WeatherCode         []int     `json:"weather_code"`
	WindDirection100m   []float32 `json:"wind_direction_100m"`
	WindDirection10m    []float32 `json:"wind_direction_10m"`
	WindDirection20m    []float32 `json:"wind_direction_20m"`
	WindDirection50m    []float32 `json:"wind_direction_50m"`
	WindDirection80m    []float32 `json:"wind_direction_80m"`
	WindGusts10m        []float32 `json:"wind_gusts_10m"`
	WindSpeed100m       []float32 `json:"wind_speed_100m"`
	WindSpeed10m        []float32 `json:"wind_speed_10m"`
	WindSpeed20m        []float32 `json:"wind_speed_20m"`
	WindSpeed50m        []float32 `json:"wind_speed_50m"`
	WindSpeed80m        []float32 `json:"wind_speed_80m"`
}

type Minutely15Units struct {
	ApparentTemperature string `json:"apparent_temperature"`
	Cape                string `json:"cape"`
	DewPoint2m          string `json:"dew_point_2m"`
	FreezingLevelHeight string `json:"freezing_level_height"`
	IsDay               string `json:"is_day"`
	LightningPotential  string `json:"lightning_potential"`
	Precipitation       string `json:"precipitation"`
	Rain                string `json:"rain"`
	RelativeHumidity2m  string `json:"relative_humidity_2m"`
	Snowfall            string `json:"snowfall"`
	SnowfallHeight      string `json:"snowfall_height"`
	SunshineDuration    string `json:"sunshine_duration"`
	Temperature2m       string `json:"temperature_2m"`
	Time                string `json:"time"`
	Visibility          string `json:"visibility"`
	WeatherCode         string `json:"weather_code"`
	WindDirection100m   string `json:"wind_direction_100m"`
	WindDirection10m    string `json:"wind_direction_10m"`
	WindDirection20m    string `json:"wind_direction_20m"`
	WindDirection50m    string `json:"wind_direction_50m"`
	WindDirection80m    string `json:"wind_direction_80m"`
	WindGusts10m        string `json:"wind_gusts_10m"`
	WindSpeed100m       string `json:"wind_speed_100m"`
	WindSpeed10m        string `json:"wind_speed_10m"`
	WindSpeed20m        string `json:"wind_speed_20m"`
	WindSpeed50m        string `json:"wind_speed_50m"`
	WindSpeed80m        string `json:"wind_speed_80m"`
}

type ForecastResponse struct {
	Current              Current         `json:"current"`
	CurrentUnits         CurrentUnits    `json:"current_units"`
	Daily                Daily           `json:"daily"`
	DailyUnits           DailyUnits      `json:"daily_units"`
	Elevation            float32         `json:"elevation"`
	GenerationtimeMs     float32         `json:"generationtime_ms"`
	Hourly               Hourly          `json:"hourly"`
	HourlyUnits          HourlyUnits     `json:"hourly_units"`
	Latitude             float32         `json:"latitude"`
	Longitude            float32         `json:"longitude"`
	Minutely15           Minutely15      `json:"minutely_15"`
	Minutely15Units      Minutely15Units `json:"minutely_15_units"`
	Timezone             string          `json:"timezone"`
	TimezoneAbbreviation string          `json:"timezone_abbreviation"`
	UtcOffsetSeconds     int             `json:"utc_offset_seconds"`
}

type ErrorForecastResponse struct {
	Error  bool   `json:"error"`
	Reason string `json:"reason"`
}

func (forecast *ForecastRequest) BuildUrl(baseUrl string, response *ForecastResponse) (string, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	u.Scheme = "https"
	v := url.Values{}
	v.Add("latitude", fmt.Sprint(forecast.Latitude))

	v.Add("longitude", fmt.Sprint(forecast.Longitude))

	if len(forecast.Hourly) > 0 {
		var hourlyParams string = ""
		for i := 0; i < len(forecast.Hourly); i++ {
			param := forecast.Hourly[i].String()
			if hourlyParams == "" && strings.TrimSpace(param) != "" {
				hourlyParams = string(param)
			} else if strings.TrimSpace(param) != "" {
				hourlyParams += "," + string(param)
			}
		}
		if strings.TrimSpace(hourlyParams) != "" {
			v.Add("hourly", hourlyParams)
		}

	}

	if len(forecast.Daily) > 0 {
		var dailyParams string = ""
		for i := 0; i < len(forecast.Daily); i++ {
			param := forecast.Daily[i].String()
			if dailyParams == "" && strings.TrimSpace(param) != "" {
				dailyParams = string(param)
			} else if strings.TrimSpace(param) != "" {
				dailyParams += "," + string(param)
			}
		}
		if strings.TrimSpace(dailyParams) != "" {
			v.Add("daily", dailyParams)
		}
	}

	if len(forecast.Current) > 0 {
		var currentParams string = ""
		for i := 0; i < len(forecast.Current); i++ {
			param := forecast.Current[i].String()
			if currentParams == "" && strings.TrimSpace(param) != "" {
				currentParams = string(param)
			} else if strings.TrimSpace(param) != "" {
				currentParams += "," + string(param)
			}
		}
		if strings.TrimSpace(currentParams) != "" {
			v.Add("current", currentParams)
		}

	}

	if len(forecast.Minutely15) > 0 {
		var minutely15Params string = ""
		for i := 0; i < len(forecast.Minutely15); i++ {
			param := forecast.Minutely15[i].String()
			if minutely15Params == "" && strings.TrimSpace(param) != "" {
				minutely15Params = string(param)
			} else if strings.TrimSpace(param) != "" {
				minutely15Params += "," + string(param)
			}
		}
		if strings.TrimSpace(minutely15Params) != "" {
			v.Add("minutely_15", minutely15Params)
		}
	}

	if forecast.Elevation > 0 {
		v.Add("elevation", fmt.Sprint(forecast.Elevation))
	}

	if strings.TrimSpace(forecast.TemperatureUnit.String()) != "" {
		v.Add("temperature_unit", forecast.TemperatureUnit.String())
	}

	if strings.TrimSpace(forecast.WindSpeedUnit.String()) != "" {
		v.Add("wind_speed_unit", forecast.WindSpeedUnit.String())
	}

	if strings.TrimSpace(forecast.PrecipitationUnit.String()) != "" {
		v.Add("precipitation_unit", forecast.PrecipitationUnit.String())
	}

	if strings.TrimSpace(forecast.Timeformat.String()) != "" {
		v.Add("timeformat", forecast.Timeformat.String())
	}

	if strings.TrimSpace(forecast.Timezone) != "" {
		v.Add("timezone", forecast.Timezone)
	}

	if forecast.PastDays > 0 {
		v.Add("past_days", fmt.Sprint(forecast.PastDays))
	}

	if forecast.ForecastDays > 0 {
		v.Add("forecast_days", fmt.Sprint(forecast.ForecastDays))
	}

	if forecast.PastHours > 0 {
		v.Add("past_hours", fmt.Sprint(forecast.PastHours))
	}

	if forecast.ForecastHours > 0 {
		v.Add("forecast_hours", fmt.Sprint(forecast.ForecastHours))
	}

	if strings.TrimSpace(forecast.StartDate) != "" {
		v.Add("start_date", fmt.Sprint(forecast.StartDate))
	}

	if strings.TrimSpace(forecast.EndDate) != "" {
		v.Add("end_date", fmt.Sprint(forecast.EndDate))
	}

	if forecast.Tilt > 0 {
		v.Add("tilt", fmt.Sprint(forecast.Tilt))
	}

	if forecast.Azimuth > 0 {
		v.Add("azimuth", fmt.Sprint(forecast.Azimuth))
	}

	if forecast.CellSelection.String() != "" {
		v.Add("cell_selection", forecast.CellSelection.String())
	}

	if strings.TrimSpace(forecast.Apikey) != "" {
		v.Add("apikey", fmt.Sprint(forecast.Apikey))
	}

	if len(forecast.Models) > 0 {
		var models string = ""
		for i := 0; i < len(forecast.Models); i++ {
			param := forecast.Models[i].String()
			if models == "" && strings.TrimSpace(param) != "" {
				models = string(param)
			} else if strings.TrimSpace(param) != "" {
				models += "," + string(param)
			}
		}
		if strings.TrimSpace(models) != "" {
			v.Add("models", fmt.Sprint(forecast.Models))
		}
	}

	url := baseUrl + "?" + v.Encode()

	fmt.Printf("Forecast request URL: %s \n", url)

	return url, nil
}

func (p ForecastParamsHourly) String() string {
	return string(p)
}

func (p ForecastParamsDaily) String() string {
	return string(p)
}

func (p ForecastParamsCurrent) String() string {
	return string(p)
}

func (p ForecastParamsMinutely15) String() string {
	return string(p)
}

func (p ForecastParamsTemperatureUnit) String() string {
	return string(p)
}

func (p ForecastParamsWindSpeedUnit) String() string {
	return string(p)
}

func (p ForecastParamsPrecipitationUnit) String() string {
	return string(p)
}

func (p ForecastParamsTimeformat) String() string {
	return string(p)
}

func (p ForecastParamsCellSelection) String() string {
	return string(p)
}

func (p ForecastParamsModels) String() string {
	return string(p)
}

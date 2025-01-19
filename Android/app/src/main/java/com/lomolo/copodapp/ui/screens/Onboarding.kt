package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowBack
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedCard
import androidx.compose.material3.OutlinedIconButton
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation

object OnboardingScreenDestination : Navigation {
    override val title = R.string.onboarding
    override val route = "onboarding"
}

@Composable
fun OnboardingScreen(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
) {
    Scaffold(topBar = {
        TopBar(
            title = {
                Text(stringResource(R.string.onboarding))
            },
            navigationIcon = {
                IconButton(
                    onClick = onGoBack,
                ) {
                    Icon(
                        Icons.AutoMirrored.TwoTone.ArrowBack,
                        contentDescription = stringResource(R.string.go_back)
                    )
                }
            },
        )
    }) { innerPadding ->
        Surface(
            modifier = modifier.fillMaxSize()
        ) {
            Column(
                Modifier
                    .fillMaxSize()
                    .padding(innerPadding)
                    .padding(8.dp)
                    .verticalScroll(
                        rememberScrollState()
                    ),
                verticalArrangement = Arrangement.spacedBy(12.dp),
            ) {
                OutlinedCard(
                    Modifier
                        .fillMaxWidth()
                        .height(360.dp)
                ) {
                    Box(
                        Modifier
                            .fillMaxSize()
                            .padding(12.dp)
                    ) {
                        Column(
                            Modifier
                                .align(Alignment.TopStart),
                        ) {
                            Icon(
                                painterResource(R.drawable.earth),
                                modifier = Modifier.size(32.dp),
                                contentDescription = stringResource(R.string.earth),
                            )
                            Spacer(Modifier.size(20.dp))
                            Text(
                                stringResource(R.string.land_title),
                                style = MaterialTheme.typography.displaySmall,
                            )
                            Text(
                                stringResource(R.string.upload_govt_issued_title)
                            )
                        }
                        Text(
                            stringResource(R.string.land_uplod_copy_text),
                            modifier = Modifier.align(Alignment.CenterStart),
                        )
                        OutlinedIconButton(
                            onClick = {},
                            modifier = Modifier.align(Alignment.BottomEnd).size(60.dp),
                        ) {
                            Icon(
                                painterResource(R.drawable.doc_paper),
                                modifier = Modifier.size(32.dp),
                                contentDescription = stringResource(R.string.upload)
                            )
                        }
                    }
                }
                OutlinedCard(
                    Modifier
                        .fillMaxWidth()
                        .height(360.dp)
                ) {
                    Box(
                        Modifier
                            .fillMaxSize()
                            .padding(12.dp)
                    ) {
                        Column(
                            Modifier
                                .align(Alignment.TopStart),
                        ) {
                            Icon(
                                painterResource(R.drawable.govt_id),
                                modifier = Modifier.size(32.dp),
                                contentDescription = stringResource(R.string.upload_govt_issued_id),
                            )
                            Spacer(Modifier.size(20.dp))
                            Text(
                                stringResource(R.string.govt_issued_id),
                                style = MaterialTheme.typography.displaySmall,
                            )
                            Text(
                                stringResource(R.string.upload_govt_issued_id)
                            )
                        }
                        Text(
                            stringResource(R.string.verify_your_id_copy_text),
                            modifier = Modifier.align(Alignment.CenterStart),
                        )
                        OutlinedIconButton(
                            onClick = {},
                            modifier = Modifier.align(Alignment.BottomEnd).size(60.dp),
                        ) {
                            Icon(
                                painterResource(R.drawable.doc_paper),
                                modifier = Modifier.size(32.dp),
                                contentDescription = stringResource(R.string.upload)
                            )
                        }
                    }
                }
                OutlinedCard(
                    Modifier
                        .fillMaxWidth()
                        .height(360.dp)
                ) {
                    Box(
                        Modifier
                            .fillMaxSize()
                            .padding(12.dp)
                    ) {
                        Column(
                            Modifier
                                .align(Alignment.TopStart),
                        ) {
                            Icon(
                                painterResource(R.drawable.account),
                                modifier = Modifier.size(32.dp),
                                contentDescription = stringResource(R.string.display_picture),
                            )
                            Spacer(Modifier.size(20.dp))
                            Text(
                                stringResource(R.string.display_picture),
                                style = MaterialTheme.typography.displaySmall,
                            )
                            Text(
                                stringResource(R.string.picture_of_you)
                            )
                        }
                        Text(
                            stringResource(R.string.upload_dp_copy_text),
                            modifier = Modifier.align(Alignment.CenterStart),
                        )
                        OutlinedIconButton(
                            onClick = {},
                            modifier = Modifier.align(Alignment.BottomEnd).size(60.dp),
                        ) {
                            Icon(
                                painterResource(R.drawable.doc_paper),
                                modifier = Modifier.size(32.dp),
                                contentDescription = stringResource(R.string.upload)
                            )
                        }
                    }
                }
            }
        }
    }
}